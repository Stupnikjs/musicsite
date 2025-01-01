package database

type DBRepo interface {
	IncrementListenCount(musicId int)
	LeaveComment(musicId int, comment string)
}
