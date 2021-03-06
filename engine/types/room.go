package types

// Room defines what we expect from room. Please notice that we do not provide
// map size here. Map size is not a requirement and if you want you can
// implement unlimited size map. Also how scores are calculated depends on
// implementation.
type Room interface {

	// Init should initialize room variables and settings. This should be
	// called before using room.
	Init() error

	// Name should return room name. It is ok if room does not have or has
	// empty name. Room name basiclly is used in multiroom server so server
	// knows where user wants to join.
	Name() string

	// SetName should set name for this room.
	SetName(name string)

	// Run should start the room.
	Run() error

	// Process should handle queued room event.
	Process() error

	// Stop should stop the room and kill everyone inside it. This is like
	// end of the world.
	Stop() error

	// AddZombie should add zombie into room.
	AddZombie(z Zombie) error

	// AddPlayer should attach client to this room.
	AddPlayer(p Player) error

	// ZombiesWon should return true if zombies won this room.
	ZombiesWon() bool

	// PlayersWon should return true if players won this room.
	PlayersWon() bool
}
