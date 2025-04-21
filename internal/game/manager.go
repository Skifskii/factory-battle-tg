package game

import (
	"errors"
	"main/internal/game/domain"
)

// type RoomManager interface {
// 	GetRoom() *domain.Room
// 	SetStatus(s domain.RoomStatus)
// 	SetRound(r int64)
// 	AddPlayer(p *domain.Player)
// 	StartSession() *domain.Player
// 	GetLeader() domain.Player
// 	GetPlayers() []domain.Player
// }

type GameManager struct {
	Rooms            map[int64]*RoomManager
	RoomIDByPlayerID map[int64]int64
}

func NewGameManager() *GameManager {
	return &GameManager{
		Rooms:            make(map[int64]*RoomManager),
		RoomIDByPlayerID: make(map[int64]int64),
	}
}

func (gm *GameManager) AddRoom(creatorID int64) (int64, error) {
	// ToDo: предусмотреть потокобезопасность
	leader := domain.NewPlayer(domain.NewUser(creatorID))
	roomID, err := gm.generateRoomID()
	if err != nil {
		return 0, err
	}

	room := domain.NewRoom(roomID, leader.ID)
	gm.Rooms[roomID] = NewRoomManager(room)

	if err := gm.AddPlayerToRoom(creatorID, roomID); err != nil {
		return 0, err
	}

	return roomID, nil
}

func (gm *GameManager) AddPlayerToRoom(userID, roomID int64) error {
	if _, exists := gm.Rooms[roomID]; !exists {
		return errors.New("room not found")
	}

	gm.RoomIDByPlayerID[userID] = roomID
	return gm.Rooms[roomID].AddPlayer(domain.NewPlayer(domain.NewUser(userID)))
}

func (gm *GameManager) StartGame(roomID int64, playerID int64, notificationChanel chan domain.Notification) error {
	rm, exists := gm.Rooms[roomID]
	if !exists {
		return errors.New("room not found")
	}

	if rm.Room.LeaderID != playerID {
		return errors.New("only leader can start game")
	}

	if len(rm.Room.Players) < 2 {
		return errors.New("a minimum of two players are required to start")
	}

	go rm.ProcessSession(notificationChanel)

	return nil
}

func (gm *GameManager) generateRoomID() (int64, error) {
	if len(gm.Rooms) == 0 {
		return 1, nil
	}

	var id int64 = 1

	for id < 100 {
		if _, exists := gm.Rooms[id]; !exists {
			return id, nil
		}
	}

	return 0, errors.New("can't generate room id")
}
