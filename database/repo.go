package database

type DBRepo interface {
	IncrementListenCount(musicId int) error
	LeaveComment(musicId int, comment string) error
	InsertSong(name string) error
}
