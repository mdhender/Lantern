//----------------------------------------------------------------------------
// ZGork - a playful imitation of Zork.
// Copyright (C) 2018 Michael D Henderson
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//----------------------------------------------------------------------------

package room

// Rooms is a map (literally) of the dungeon.
type Rooms map[Key]Room

// Room defines the structure of a room.
type Room struct {
	Name string
	Key  Key
	Desc string

	Exits struct {
		Frak, Cross, Down, East, Enter, Exit, Land, Launc, Leave, Ne, North, Nw, Out, Southeast, South, Sw, Up, West Key
	}
}

// AddExit adds an exit to a room
func (r Room) AddExit(e Exit) Room {
	switch e.Dir {
	case HASHBANGHASHBANGHASH:
		r.Exits.Frak = e.Target
	case CROSS:
		r.Exits.Cross = e.Target
	case DOWN:
		r.Exits.Down = e.Target
	case EAST:
		r.Exits.East = e.Target
	case ENTER:
		r.Exits.Enter = e.Target
	case EXIT:
		r.Exits.Exit = e.Target
	case LAND:
		r.Exits.Land = e.Target
	case LAUNC:
		r.Exits.Launc = e.Target
	case LEAVE:
		r.Exits.Leave = e.Target
	case NE:
		r.Exits.Ne = e.Target
	case NORTH:
		r.Exits.North = e.Target
	case NW:
		r.Exits.Nw = e.Target
	case OUT:
		r.Exits.Out = e.Target
	case SE:
		r.Exits.Southeast = e.Target
	case SOUTH:
		r.Exits.South = e.Target
	case SW:
		r.Exits.Sw = e.Target
	case UP:
		r.Exits.Up = e.Target
	case WEST:
		r.Exits.West = e.Target
	}
	return r
}

// NewMap creates a new map of rooms
func NewMap() Rooms {
	rooms := make(map[Key]Room)
	for _, r := range AllRooms() {
		rooms[r.Key] = r
	}

	for _, e := range AllExits() {
		rooms[e.Source] = rooms[e.Source].AddExit(e)
	}

	return rooms
}

// AllRooms creates all the rooms
func AllRooms() []Room {
	return []Room{
		Room{
			Name: "West of House",
			Key:  WHOUS,
			Desc: "This is an open field west of a white house, with a boarded front door.",
		},
		Room{
			Name: "North of House",
			Key:  NHOUS,
			Desc: "You are facing the north side of a white house.  There is no door here,\nand all the windows are barred.",
		},
		Room{
			Name: "South of House",
			Key:  SHOUS,
			Desc: "You are facing the south side of a white house. There is no door here,\nand all the windows are barred.",
		},
		Room{
			Name: "Behind House",
			Key:  EHOUS,
		},
		Room{
			Name: "Kitchen",
			Key:  KITCH,
		},
		Room{
			Name: "Attic",
			Key:  ATTIC,
			Desc: "This is the attic.  The only exit is stairs that lead down.",
		},
		Room{
			Name: "Living Room",
			Key:  LROOM,
		},
		Room{
			Name: "Forest",
			Key:  FORE1,
			Desc: "This is a forest, with trees in all directions around you.",
		},
		Room{
			Name: "Forest",
			Key:  FORE2,
			Desc: "This is a dimly lit forest, with large trees all around.  To the\neast, there appears to be sunlight.",
		},
		Room{
			Name: "Forest",
			Key:  FORE3,
			Desc: "This is a dimly lit forest, with large trees all around.  One\nparticularly large tree with some low branches stands here.",
		},
		Room{
			Name: "Up a Tree",
			Key:  TREE,
		},
		Room{
			Name: "Forest",
			Key:  FORE4,
			Desc: "This is a large forest, with trees obstructing all views except\nto the east, where a small clearing may be seen through the trees.",
		},
		Room{
			Name: "Forest",
			Key:  FORE5,
			Desc: "This is a forest, with trees in all directions around you.",
		},
		Room{
			Name: "Clearing",
			Key:  CLEAR,
		},
		Room{
			Name: "Cellar",
			Key:  CELLA,
		},
		Room{
			Name: "The Troll Room",
			Key:  MTROL,
			Desc: "This is a small room with passages off in all directions. \nBloodstains and deep scratches (perhaps made by an axe) mar the\nwalls.",
		},
		Room{
			Name: "Studio",
			Key:  STUDI,
			Desc: "This is what appears to have been an artist's studio.  The walls\nand floors are splattered with paints of 69 different colors. \nStrangely enough, nothing of value is hanging here.  At the north and\nnorthwest of the room are open doors (also covered with paint).  An\nextremely dark and narrow chimney leads up from a fireplace; although\nyou might be able to get up it, it seems unlikely you could get back\ndown.",
		},
		Room{
			Name: "Gallery",
			Key:  GALLE,
			Desc: "This is an art gallery.  Most of the paintings which were here\nhave been stolen by vandals with exceptional taste.  The vandals\nleft through either the north, south, or west exits.",
		},
		Room{
			Name: "Maze",
			Key:  MAZE1,
			Desc: "This is part of a maze of twisty little passages, all alike.",
		},
		Room{
			Name: "Maze",
			Key:  MAZE2,
			Desc: "This is part of a maze of twisty little passages, all alike.",
		},
		Room{
			Name: "Maze",
			Key:  MAZE3,
			Desc: "This is part of a maze of twisty little passages, all alike.",
		},
		Room{
			Name: "Maze",
			Key:  MAZE4,
			Desc: "This is part of a maze of twisty little passages, all alike.",
		},
		Room{
			Name: "You have come to a dead end in the maze.",
			Key:  DEAD1,
			Desc: "Dead End",
		},
		Room{
			Name: "Maze",
			Key:  MAZE5,
			Desc: "This is part of a maze of twisty little passages, all alike.",
		},
		Room{
			Name: "You have come to a dead end in the maze.",
			Key:  DEAD2,
			Desc: "Dead End",
		},
		Room{
			Name: "Maze",
			Key:  MAZE6,
			Desc: "This is part of a maze of twisty little passages, all alike.",
		},
		Room{
			Name: "Maze",
			Key:  MAZE7,
			Desc: "This is part of a maze of twisty little passages, all alike.",
		},
		Room{
			Name: "Maze",
			Key:  MAZE8,
			Desc: "This is part of a maze of twisty little passages, all alike.",
		},
		Room{
			Name: "Dead End",
			Key:  DEAD3,
			Desc: "Dead End",
		},
		Room{
			Name: "Maze",
			Key:  MAZE9,
			Desc: "This is part of a maze of twisty little passages, all alike.",
		},
		Room{
			Name: "Maze",
			Key:  MAZ10,
			Desc: "This is part of a maze of twisty little passages, all alike.",
		},
		Room{
			Name: "Maze",
			Key:  MAZ11,
			Desc: "This is part of a maze of twisty little passages, all alike.",
		},
		Room{
			Name: "Grating Room",
			Key:  MGRAT,
		},
		Room{
			Name: "Maze",
			Key:  MAZ12,
			Desc: "This is part of a maze of twisty little passages, all alike.",
		},
		Room{
			Name: "Dead End",
			Key:  DEAD4,
			Desc: "Dead End",
		},
		Room{
			Name: "Maze",
			Key:  MAZ13,
			Desc: "This is part of a maze of twisty little passages, all alike.",
		},
		Room{
			Name: "Maze",
			Key:  MAZ14,
			Desc: "This is part of a maze of twisty little passages, all alike.",
		},
		Room{
			Name: "Maze",
			Key:  MAZ15,
			Desc: "This is part of a maze of twisty little passages, all alike.",
		},
		Room{
			Name: "Cyclops Room",
			Key:  CYCLO,
		},
		Room{
			Name: "Strange Passage",
			Key:  BLROO,
			Desc: "This is a long passage.  To the south is one entrance.  On the\neast there is an old wooden door, with a large hole in it (about\ncyclops sized).",
		},
		Room{
			Name: "Treasure Room",
			Key:  TREAS,
			Desc: "This is a large room, whose north wall is solid granite.  A number\nof discarded bags, which crumble at your touch, are scattered about\non the floor.  There is an exit down and what appears to be a newly\ncreated passage to the east.",
		},
		Room{
			Name: "Deep Ravine",
			Key:  RAVI1,
			Desc: "This is a deep ravine at a crossing with an east-west crawlway. \nSome stone steps are at the south of the ravine and a steep staircase\ndescends.",
		},
		Room{
			Name: "Rocky Crawl",
			Key:  CRAW1,
			Desc: "This is a crawlway with a three-foot high ceiling.  Your footing\nis very unsure here due to the assortment of rocks underfoot. \nPassages can be seen in the east, west, and northwest corners of the\npassage.",
		},
		Room{
			Name: "Reservoir South",
			Key:  RESES,
		},
		Room{
			Name: "Reservoir",
			Key:  RESER,
		},
		Room{
			Name: "Reservoir North",
			Key:  RESEN,
		},
		Room{
			Name: "Stream View",
			Key:  STREA,
			Desc: "You are standing on a path beside a gently flowing stream.  The path\ntravels to the north and the east.",
		},
		Room{
			Name: "Stream",
			Key:  INSTR,
			Desc: "You are on the gently flowing stream.  The upstream route is too narrow\nto  navigate and the downstream route is invisible due to twisting\nwalls.  There is a narrow beach to land on.",
		},
		Room{
			Name: "Egyptian Room",
			Key:  EGYPT,
			Desc: "This is a room which looks like an Egyptian tomb.  There is an\nascending staircase in the room as well as doors, east and south.",
		},
		Room{
			Name: "Glacier Room",
			Key:  ICY,
		},
		Room{
			Name: "Ruby Room",
			Key:  RUBYR,
			Desc: "This is a small chamber behind the remains of the Great Glacier.\nTo the south and west are small passageways.",
		},
		Room{
			Name: "Atlantis Room",
			Key:  ATLAN,
			Desc: "This is an ancient room, long under water.  There are exits here\nto the southeast and upward.",
		},
		Room{
			Name: "Deep Canyon",
			Key:  CANY1,
			Desc: "You are on the south edge of a deep canyon.  Passages lead off\nto the east, south, and northwest.  You can hear the sound of\nflowing water below.",
		},
		Room{
			Name: "Loud Room",
			Key:  ECHO,
			Desc: "This is a large room with a ceiling which cannot be detected from\nthe ground. There is a narrow passage from east to west and a stone\nstairway leading upward.  The room is extremely noisy.  In fact, it is\ndifficult to hear yourself think.",
		},
		Room{
			Name: "Mirror Room",
			Key:  MIRR1,
		},
		Room{
			Name: "Mirror Room",
			Key:  MIRR2,
		},
		Room{
			Name: "Cave",
			Key:  CAVE1,
			Desc: "This is a small cave with an entrance to the north and a stairway\nleading down.",
		},
		Room{
			Name: "Cave",
			Key:  CAVE2,
			Desc: "This is a tiny cave with entrances west and north, and a dark,\nforbidding staircase leading down.",
		},
		Room{
			Name: "Steep Crawlway",
			Key:  CRAW2,
			Desc: "This is a steep and narrow crawlway.  There are two exits nearby to\nthe south and southwest.",
		},
		Room{
			Name: "Narrow Crawlway",
			Key:  CRAW3,
			Desc: "This is a narrow crawlway.  The crawlway leads from north to south.\nHowever the south passage divides to the south and southwest.",
		},
		Room{
			Name: "Cold Passage",
			Key:  PASS3,
			Desc: "This is a cold and damp corridor where a long east-west passageway\nintersects with a northward path.",
		},
		Room{
			Name: "Winding Passage",
			Key:  PASS4,
			Desc: "This is a winding passage.  It seems that there is only an exit\non the east end although the whirring from the round room can be\nheard faintly to the north.",
		},
		Room{
			Name: "Mine Entrance",
			Key:  ENTRA,
			Desc: "You are standing at the entrance of what might have been a coal\nmine. To the northeast and the northwest are entrances to the mine,\nand there is another exit on the south end of the room.",
		},
		Room{
			Name: "Squeaky Room",
			Key:  SQUEE,
			Desc: "You are a small room.  Strange squeaky sounds may be heard coming from\nthe passage at the west end.  You may also escape to the south.",
		},
		Room{
			Name: "Shaft Room",
			Key:  TSHAF,
			Desc: "This is a large room, in the middle of which is a small shaft\ndescending through the floor into darkness below.  To the west and\nthe north are exits from this room.  Constructed over the top of the\nshaft is a metal framework to which a heavy iron chain is attached.",
		},
		Room{
			Name: "Wooden Tunnel",
			Key:  TUNNE,
			Desc: "This is a narrow tunnel with large wooden beams running across\nthe ceiling and around the walls.  A path from the south splits into\npaths running west and northeast.",
		},
		Room{
			Name: "Smelly Room",
			Key:  SMELL,
			Desc: "This is a small non-descript room.  However, from the direction\nof a small descending staircase a foul odor can be detected.  To the\neast is a narrow path.",
		},
		Room{
			Name: "Gas Room",
			Key:  BOOM,
			Desc: "This is a small room which smells strongly of coal gas.",
		},
		Room{
			Name: "Ladder Top",
			Key:  TLADD,
			Desc: "This is a very small room.  In the corner is a rickety wooden\nladder, leading downward.  It might be safe to descend.  There is\nalso a staircase leading upward.",
		},
		Room{
			Name: "Ladder Bottom",
			Key:  BLADD,
			Desc: "This is a rather wide room.  On one side is the bottom of a\nnarrow wooden ladder.  To the northeast and the south are passages\nleaving the room.",
		},
		Room{
			Name: "Dead End",
			Key:  DEAD7,
			Desc: "Dead End",
		},
		Room{
			Name: "Timber Room",
			Key:  TIMBE,
			Desc: "This is a long and narrow passage, which is cluttered with broken\ntimbers.  A wide passage comes from the north and turns at the \nsouthwest corner of the room into a very narrow passageway.",
		},
		Room{
			Name: "Lower Shaft",
			Key:  BSHAF,
			Desc: "This is a small square room which is at the bottom of a long\nshaft. To the east is a passageway and to the northeast a very narrow\npassage. In the shaft can be seen a heavy iron chain.",
		},
		Room{
			Name: "Machine Room",
			Key:  MACHI,
		},
		Room{
			Name: "Bat Room",
			Key:  BATS,
		},
		Room{
			Name: "Coal mine",
			Key:  MINE1,
			Desc: "This is a non-descript part of a coal mine.",
		},
		Room{
			Name: "Coal mine",
			Key:  MINE2,
			Desc: "This is a non-descript part of a coal mine.",
		},
		Room{
			Name: "Coal mine",
			Key:  MINE3,
			Desc: "This is a non-descript part of a coal mine.",
		},
		Room{
			Name: "Coal mine",
			Key:  MINE4,
			Desc: "This is a non-descript part of a coal mine.",
		},
		Room{
			Name: "Coal mine",
			Key:  MINE5,
			Desc: "This is a non-descript part of a coal mine.",
		},
		Room{
			Name: "Coal mine",
			Key:  MINE6,
			Desc: "This is a non-descript part of a coal mine.",
		},
		Room{
			Name: "Coal mine",
			Key:  MINE7,
			Desc: "This is a non-descript part of a coal mine.",
		},
		Room{
			Name: "Dome Room",
			Key:  DOME,
		},
		Room{
			Name: "Torch Room",
			Key:  MTORC,
		},
		Room{
			Name: "North-South Crawlway",
			Key:  CRAW4,
			Desc: "This is a north-south crawlway; a passage goes to the east also.\nThere is a hole above, but it provides no opportunities for climbing.",
		},
		Room{
			Name: "West of Chasm",
			Key:  CHAS2,
			Desc: "You are on the west edge of a chasm, the bottom of which cannot be\nseen. The east side is sheer rock, providing no exits.  A narrow\npassage goes west, and the path you are on continues to the north and\nsouth.",
		},
		Room{
			Name: "East-West Passage",
			Key:  PASS1,
			Desc: "This is a narrow east-west passageway.  There is a narrow stairway\nleading down at the north end of the room.",
		},
		Room{
			Name: "Round room",
			Key:  CAROU,
		},
		Room{
			Name: "North-South Passage",
			Key:  PASS5,
			Desc: "This is a high north-south passage, which forks to the northeast.",
		},
		Room{
			Name: "Chasm",
			Key:  CHAS1,
			Desc: "A chasm runs southwest to northeast.  You are on the south edge; the\npath exits to the south and to the east.",
		},
		Room{
			Name: "Damp Cave",
			Key:  CAVE3,
			Desc: "This is a cave.  Passages exit to the south and to the east, but\nthe cave narrows to a crack to the west.  The earth is particularly\ndamp here.",
		},
		Room{
			Name: "Ancient Chasm",
			Key:  CHAS3,
			Desc: "A chasm, evidently produced by an ancient river, runs through the\ncave here.  Passages lead off in all directions.",
		},
		Room{
			Name: "Dead End",
			Key:  DEAD5,
			Desc: "Dead End",
		},
		Room{
			Name: "Dead End",
			Key:  DEAD6,
			Desc: "Dead End",
		},
		Room{
			Name: "Engravings Cave",
			Key:  CAVE4,
			Desc: "You have entered a cave with passages leading north and southeast.",
		},
		Room{
			Name: "Riddle Room",
			Key:  RIDDL,
			Desc: "This is a room which is bare on all sides.  There is an exit down. \nTo the east is a great door made of stone.  Above the stone, the\nfollowing words are written: 'No man shall enter this room without\nsolving this riddle:\n\n  What is tall as a house,\n\t  round as a cup, \n\t  and all the king's horses can't draw it up?'\n\n(Reply via 'ANSWER 'answer'')",
		},
		Room{
			Name: "Pearl Room",
			Key:  MPEAR,
			Desc: "This is a former broom closet.  The exits are to the east and west.",
		},
		Room{
			Name: "Entrance to Hades",
			Key:  LLD1,
		},
		Room{
			Name: "Land of the Living Dead",
			Key:  LLD2,
		},
		Room{
			Name: "Grail Room",
			Key:  MGRAI,
			Desc: "You are standing in a small circular room with a pedestal.  A set of\nstairs leads up, and passages leave to the east and west.",
		},
		Room{
			Name: "Temple",
			Key:  TEMP1,
			Desc: "This is the west end of a large temple.  On the south wall is an \nancient inscription, probably a prayer in a long-forgotten language. \nThe north wall is solid granite.  The entrance at the west end of the\nroom is through huge marble pillars.",
		},
		Room{
			Name: "Altar",
			Key:  TEMP2,
			Desc: "This is the east end of a large temple.  In front of you is what\nappears to be an altar.",
		},
		Room{
			Name: "Dam",
			Key:  DAM,
		},
		Room{
			Name: "Dam Lobby",
			Key:  LOBBY,
			Desc: "This room appears to have been the waiting room for groups touring\nthe dam.  There are exits here to the north and east marked\n'Private', though the doors are open, and an exit to the south.",
		},
		Room{
			Name: "Maintenance Room",
			Key:  MAINT,
			Desc: "This is what appears to have been the maintenance room for Flood\nControl Dam #3, judging by the assortment of tool chests around the\nroom.  Apparently, this room has been ransacked recently, for most of\nthe valuable equipment is gone. On the wall in front of you is a\ngroup of buttons, which are labelled in EBCDIC. However, they are of\ndifferent colors:  Blue, Yellow, Brown, and Red. The doors to this\nroom are in the west and south ends.",
		},
		Room{
			Name: "Dam Base",
			Key:  DOCK,
			Desc: "You are at the base of Flood Control Dam #3, which looms above you\nand to the north.  The river Frigid is flowing by here.  Across the\nriver are the White Cliffs which seem to form a giant wall stretching\nfrom north to south along the east shore of the river as it winds its\nway downstream.",
		},
		Room{
			Name: "Frigid River",
			Key:  RIVR1,
			Desc: "You are on the River Frigid in the vicinity of the Dam.  The river\nflows quietly here.  There is a landing on the west shore.",
		},
		Room{
			Name: "Frigid River",
			Key:  RIVR2,
			Desc: "The River turns a corner here making it impossible to see the\nDam.  The White Cliffs loom on the east bank and large rocks prevent\nlanding on the west.",
		},
		Room{
			Name: "Frigid River",
			Key:  RIVR3,
			Desc: "The river descends here into a valley.  There is a narrow beach on\nthe east below the cliffs and there is some shore on the west which\nmay be suitable.  In the distance a faint rumbling can be heard.",
		},
		Room{
			Name: "White Cliffs Beach",
			Key:  WCLF1,
			Desc: "You are on a narrow strip of beach which runs along the base of the\nWhite Cliffs. The only path here is a narrow one, heading south\nalong the Cliffs.",
		},
		Room{
			Name: "White Cliffs Beach",
			Key:  WCLF2,
			Desc: "You are on a rocky, narrow strip of beach beside the Cliffs.  A\nnarrow path leads north along the shore.",
		},
		Room{
			Name: "Frigid River",
			Key:  RIVR4,
			Desc: "The river is running faster here and the sound ahead appears to be\nthat of rushing water.  On the west shore is a sandy beach.  A small\narea of beach can also be seen below the Cliffs.",
		},
		Room{
			Name: "Frigid River",
			Key:  RIVR5,
			Desc: "The sound of rushing water is nearly unbearable here.  On the west\nshore is a large landing area.",
		},
		Room{
			Name: "Moby lossage",
			Key:  FCHMP,
		},
		Room{
			Name: "Shore",
			Key:  FANTE,
			Desc: "You are on the shore of the River.  The river here seems somewhat\ntreacherous.  A path travels from north to south here, the south end\nquickly turning around a sharp corner.",
		},
		Room{
			Name: "Sandy Beach",
			Key:  BEACH,
			Desc: "You are on a large sandy beach at the shore of the river, which is\nflowing quickly by.  A path runs beside the river to the south here.",
		},
		Room{
			Name: "Rocky Shore",
			Key:  RCAVE,
			Desc: "You are on the west shore of the river.  An entrance to a cave is\nto the northwest.  The shore is very rocky here.",
		},
		Room{
			Name: "Small Cave",
			Key:  TCAVE,
			Desc: "This is a small cave whose exits are on the south and northwest.",
		},
		Room{
			Name: "Aragain Falls",
			Key:  FALLS,
		},
		Room{
			Name: "Rainbow Room",
			Key:  RAINB,
			Desc: "You are on top of a rainbow (I bet you never thought you would walk\non a rainbow), with a magnificent view of the Falls.  The rainbow\ntravels east-west here.  There is an NBC Commissary here.",
		},
		Room{
			Name: "End of Rainbow",
			Key:  POG,
			Desc: "You are on a small, rocky beach on the continuation of the Frigid\nRiver past the Falls.  The beach is narrow due to the presence of the\nWhite Cliffs.  The river canyon opens here and sunlight shines in\nfrom above. A rainbow crosses over the falls to the west and a narrow\npath continues to the southeast.",
		},
		Room{
			Name: "Canyon Bottom",
			Key:  CLBOT,
			Desc: "You are beneath the walls of the river canyon which may be climbable\nhere.  There is a small stream here, which is the lesser part of the\nrunoff of Aragain Falls. To the north is a narrow path.",
		},
		Room{
			Name: "Rocky Ledge",
			Key:  CLMID,
			Desc: "You are on a ledge about halfway up the wall of the river canyon.\nYou can see from here that the main flow from Aragain Falls twists\nalong a passage which it is impossible to enter.  Below you is the\ncanyon bottom.  Above you is more cliff, which still appears\nclimbable.",
		},
		Room{
			Name: "Canyon View",
			Key:  CLTOP,
			Desc: "You are at the top of the Great Canyon on its south wall.  From here\nthere is a marvelous view of the Canyon and parts of the Frigid River\nupstream.  Across the canyon, the walls of the White Cliffs still\nappear to loom far above.  Following the Canyon upstream (north and\nnorthwest), Aragain Falls may be seen, complete with rainbow. \nFortunately, my vision is better than average and I can discern the\ntop of the Flood Control Dam #3 far to the distant north.  To the\nwest and south can be seen an immense forest, stretching for miles\naround.  It is possible to climb down into the canyon from here.",
		},
		Room{
			Name: "Volcano Bottom",
			Key:  VLBOT,
			Desc: "You are at the bottom of a large dormant volcano.  High above you\nlight may be seen entering from the cone of the volcano.  The only\nexit here is to the north.",
		},
		Room{
			Name: "Volcano Core",
			Key:  VAIR1,
			Desc: "You are about one hundred feet above the bottom of the volcano.  The\ntop of the volcano is clearly visible here.",
		},
		Room{
			Name: "Volcano near small ledge",
			Key:  VAIR2,
			Desc: "You are about two hundred feet above the volcano floor.  Looming\nabove is the rim of the volcano.  There is a small ledge on the west\nside.",
		},
		Room{
			Name: "Volcano near viewing ledge",
			Key:  VAIR3,
			Desc: "You are high above the floor of the volcano.  From here the rim of\nthe volcano looks very narrow and you are very near it.  To the \neast is what appears to be a viewing ledge, too thin to land on.",
		},
		Room{
			Name: "Volcano near wide ledge",
			Key:  VAIR4,
			Desc: "You are near the rim of the volcano which is only about 15 feet\nacross.  To the west, there is a place to land on a wide ledge.",
		},
		Room{
			Name: "Narrow Ledge",
			Key:  LEDG2,
			Desc: "You are on a narrow ledge overlooking the inside of an old dormant\nvolcano.  This ledge appears to be about in the middle between the\nfloor below and the rim above. There is an exit here to the south.",
		},
		Room{
			Name: "Library",
			Key:  LIBRA,
			Desc: "This is a room which must have been a large library, probably\nfor the royal family.  All of the shelves appear to have been gnawed\nto pieces by unfriendly gnomes.  To the north is an exit.",
		},
		Room{
			Name: "Volcano View",
			Key:  LEDG3,
			Desc: "You are on a ledge in the middle of a large volcano.  Below you\nthe volcano bottom can be seen and above is the rim of the volcano.\nA couple of ledges can be seen on the other side of the volcano;\nit appears that this ledge is intermediate in elevation between\nthose on the other side.  The exit from this room is to the east.",
		},
		Room{
			Name: "Wide Ledge",
			Key:  LEDG4,
		},
		Room{
			Name: "Dusty Room",
			Key:  SAFE,
		},
		Room{
			Name: "Lava Room",
			Key:  LAVA,
			Desc: "This is a small room, whose walls are formed by an old lava flow.\nThere are exits here to the west and the south.",
		},
		Room{
			Name: "Low Room",
			Key:  MAGNE,
		},
		Room{
			Name: "Machine Room",
			Key:  CMACH,
		},
		Room{
			Name: "Dingy Closet",
			Key:  CAGER,
			Desc: "This is a dingy closet adjacent to the machine room.  On one wall\nis a small sticker which says\n\t\tProtected by\n\t\t  FROBOZZ\n\t     Magic Alarm Company\n\t      (Hello, footpad!)\n",
		},
		Room{
			Name: "Cage",
			Key:  CAGED,
			Desc: "You are trapped inside a steel cage.",
		},
		Room{
			Name: "Top of Well",
			Key:  TWELL,
			Desc: "You are at the top of the well.  Well done.  There are etchings on\nthe side of the well. There is a small crack across the floor at the\nentrance to a room on the east, but it can be crossed easily.",
		},
		Room{
			Name: "Circular Room",
			Key:  BWELL,
			Desc: "This is a damp circular room, whose walls are made of brick and\nmortar.  The roof of this room is not visible, but there appear to be\nsome etchings on the walls.  There is a passageway to the west.",
		},
		Room{
			Name: "Tea Room",
			Key:  ALICE,
			Desc: "This is a small square room, in the center of which is a large\noblong table, no doubt set for afternoon tea.  It is clear from the\nobjects on the table that the users were indeed mad.  In the eastern\ncorner of the room is a small hole (no more than four inches high). \nThere are passageways leading away to the west and the northwest.",
		},
		Room{
			Name: "Posts Room",
			Key:  ALISM,
			Desc: "This is an enormous room, in the center of which are four wooden\nposts delineating a rectangular area, above which is what appears to\nbe a wooden roof.  In fact, all objects in this room appear to be\nabnormally large. To the east is a passageway.  There is a large\nchasm on the west and the northwest.",
		},
		Room{
			Name: "Pool Room",
			Key:  ALITR,
			Desc: "This is a large room, one half of which is depressed.  There is a\nlarge leak in the ceiling through which brown colored goop is\nfalling.  The only exit to this room is to the west.",
		},
		Room{
			Name: "Bank Entrance",
			Key:  BKENT,
			Desc: "This is the large entrance hall of the Bank of Zork, the largest\nbanking institution of the Great Underground Empire. A partial\naccount of its history is in 'The Lives of the Twelve Flatheads' with\nthe chapter on J. Pierpont Flathead.  A more detailed history (albeit\nless objective) may be found in Flathead's outrageous autobiography\n'I'm Rich and You Aren't - So There!'.\nMost of the furniture has been ravaged by passing scavengers.  All\nthat remains are two signs at the Northwest and Northeast corners of\nthe room, which say\n  \n      <--  WEST VIEWING ROOM        EAST VIEWING ROOM  -->  \n",
		},
		Room{
			Name: "West Teller's Room",
			Key:  BKTW,
		},
		Room{
			Name: "East Teller's Room",
			Key:  BKTE,
		},
		Room{
			Name: "Viewing Room",
			Key:  BKVW,
			Desc: "This is a room used by holders of safety deposit boxes to view\ntheir contents.  On the north side of the room is a sign which says \n\t\n   REMAIN HERE WHILE THE BANK OFFICER RETRIEVES YOUR DEPOSIT BOX\n    WHEN YOU ARE FINISHED, LEAVE THE BOX, AND EXIT TO THE SOUTH  \n     AN ADVANCED PROTECTIVE DEVICE PREVENTS ALL CUSTOMERS FROM\n      REMOVING ANY SAFETY DEPOSIT BOX FROM THIS VIEWING AREA!\n               Thank You for banking at the Zork!\n",
		},
		Room{
			Name: "Viewing Room",
			Key:  BKVE,
			Desc: "This is a room used by holders of safety deposit boxes to view\ntheir contents.  On the north side of the room is a sign which says \n\t\n   REMAIN HERE WHILE THE BANK OFFICER RETRIEVES YOUR DEPOSIT BOX\n    WHEN YOU ARE FINISHED, LEAVE THE BOX, AND EXIT TO THE SOUTH  \n     AN ADVANCED PROTECTIVE DEVICE PREVENTS ALL CUSTOMERS FROM\n      REMOVING ANY SAFETY DEPOSIT BOX FROM THIS VIEWING AREA!\n               Thank You for banking at the Zork!\n",
		},
		Room{
			Name: "Small Room",
			Key:  BKTWI,
			Desc: "This is a small, bare room with no distinguishing features. There\nare no exits from this room.",
		},
		Room{
			Name: "Vault",
			Key:  BKVAU,
			Desc: "This is the Vault of the Bank of Zork, in which there are no doors.",
		},
		Room{
			Name: "Safety Depository",
			Key:  BKBOX,
			Desc: "This is a large rectangular room.  The east and west walls here\nwere used for storing safety deposit boxes.  As might be expected,\nall have been carefully removed by evil persons.  To the east, west,\nand south of the room are large doorways. The northern 'wall'\nof the room is a shimmering curtain of light.  In the center of the\nroom is a large stone cube, about 10 feet on a side.  Engraved on \nthe side of the cube is some lettering.",
		},
		Room{
			Name: "Chairman's Office",
			Key:  BKEXE,
			Desc: "This room was the office of the Chairman of the Bank of Zork.\nLike the other rooms here, it has been extensively vandalized.\nThe lone exit is to the north.",
		},
		Room{
			Name: "Small Square Room",
			Key:  CPANT,
			Desc: "This is a small square room, in the middle of which is a recently \ncreated hole through which you can barely discern the floor some ten\nfeet below.  It doesn't seem likely you could climb back up.  There\nare exits to the west and south.",
		},
		Room{
			Name: "Side Room",
			Key:  CPOUT,
		},
		Room{
			Name: "Room in a Puzzle",
			Key:  CP,
		},
		Room{
			Name: "Dreary Room",
			Key:  PALAN,
		},
		Room{
			Name: "Tiny Room",
			Key:  PRM,
		},
		Room{
			Name: "Slide Room",
			Key:  SLIDE,
		},
		Room{
			Name: "Slide",
			Key:  SLID1,
			Desc: "This is an uncomfortable spot within the coal chute.  The rope to\nwhich you are clinging can be seen rising into the darkness above.\nThere is more rope dangling below you.",
		},
		Room{
			Name: "Slide",
			Key:  SLID2,
			Desc: "This is another spot within the coal chute.  Above you the rope\nclimbs into darkness and the end of the rope is dangling five feet\nbeneath you.",
		},
		Room{
			Name: "Slide",
			Key:  SLID3,
			Desc: "You have reached the end of your rope.  Below you is darkness as\nthe chute makes a sharp turn.  On the east here is a small ledge\nwhich you might be able to stand on.",
		},
		Room{
			Name: "Slide Ledge",
			Key:  SLEDG,
			Desc: "This is a narrow ledge abutting the coal chute, in which a rope can\nbe seen passing downward.  Behind you, to the south, is a small room.",
		},
		Room{
			Name: "Sooty Room",
			Key:  SPAL,
			Desc: "This is a small room with rough walls, and a ceiling which is steeply\nsloping from north to south. There is coal dust covering almost\neverything, and little bits of coal are scattered around the only exit\n(which is a narrow passage to the north). In one corner of the room is\nan old coal stove which lights the room with a cheery red glow.  There\nis a very narrow crack in the north wall.",
		},
		Room{
			Name: "Hallway",
			Key:  MRD,
		},
		Room{
			Name: "Hallway",
			Key:  MRG,
		},
		Room{
			Name: "Hallway",
			Key:  MRC,
		},
		Room{
			Name: "Hallway",
			Key:  MRB,
		},
		Room{
			Name: "Hallway",
			Key:  MRA,
		},
		Room{
			Name: "Narrow Room",
			Key:  MRDE,
		},
		Room{
			Name: "Narrow Room",
			Key:  MRDW,
		},
		Room{
			Name: "Narrow Room",
			Key:  MRGE,
		},
		Room{
			Name: "Narrow Room",
			Key:  MRGW,
		},
		Room{
			Name: "Narrow Room",
			Key:  MRCE,
		},
		Room{
			Name: "Narrow Room",
			Key:  MRCW,
		},
		Room{
			Name: "Narrow Room",
			Key:  MRBE,
		},
		Room{
			Name: "Narrow Room",
			Key:  MRBW,
		},
		Room{
			Name: "Narrow Room",
			Key:  MRAE,
		},
		Room{
			Name: "Narrow Room",
			Key:  MRAW,
		},
		Room{
			Name: "Inside Mirror",
			Key:  INMIR,
		},
		Room{
			Name: "Stone Room",
			Key:  MRANT,
			Desc: "You are standing near one end of a long, dimly lit hall.  At the\nsouth stone stairs ascend.  To the north the corridor is illuminated\nby torches set high in the walls, out of reach.  On one wall is a red\nbutton.",
		},
		Room{
			Name: "Small Room",
			Key:  MREYE,
		},
		Room{
			Name: "Tomb of the Unknown Implementer",
			Key:  TOMB,
		},
		Room{
			Name: "Crypt",
			Key:  CRYPT,
		},
		Room{
			Name: "Top of Stairs",
			Key:  TSTRS,
			Desc: "You are standing at the top of a flight of stairs that lead down to\na passage below.  Dim light, as from torches, can be seen in the\npassage.  Behind you the stairs lead into untouched rock.",
		},
		Room{
			Name: "East Corridor",
			Key:  ECORR,
			Desc: "This is a corridor with polished marble walls.  The corridor\nwidens into larger areas as it turns west at its northern and\nsouthern ends.",
		},
		Room{
			Name: "West Corridor",
			Key:  WCORR,
			Desc: "This is a corridor with polished marble walls.  The corridor\nwidens into larger areas as it turns east at its northern and\nsouthern ends.",
		},
		Room{
			Name: "South Corridor",
			Key:  SCORR,
		},
		Room{
			Name: "Narrow Corridor",
			Key:  BDOOR,
		},
		Room{
			Name: "Dungeon Entrance",
			Key:  FDOOR,
		},
		Room{
			Name: "North Corridor",
			Key:  NCORR,
		},
		Room{
			Name: "Parapet",
			Key:  PARAP,
		},
		Room{
			Name: "Prison Cell",
			Key:  CELL,
		},
		Room{
			Name: "Prison Cell",
			Key:  PCELL,
		},
		Room{
			Name: "Prison Cell",
			Key:  NCELL,
		},
		Room{
			Name: "Treasury of Zork",
			Key:  NIRVA,
			Desc: "This is a room of large size, richly appointed and decorated\nin a style that bespeaks exquisite taste.  To judge from its\ncontents, it is the ultimate storehouse of the treasures of Zork.\n     There are chests here containing precious jewels, mountains of\nzorkmids, rare paintings, ancient statuary, and beguiling curios.\n     In one corner of the room is a bookcase boasting such volumes as\n'The History of the Great Underground Empire,' 'The Lives of the\nTwelve Flatheads,' 'The Wisdom of the Implementors,' and other\ninformative and inspiring works.\n     On one wall is a completely annotated map of the Great Underground\nEmpire, showing points of interest, various troves of treasure, and\nindicating the locations of several superior scenic views.\n     On a desk at the far end of the room may be found stock\ncertificates representing a controlling interest in FrobozzCo\nInternational, the multinational conglomerate and parent company of\nthe Frobozz Magic Boat Co., etc.",
		},
	}
}
