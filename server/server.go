package server

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"go-dart/common"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

type Server struct {
	games map[int]Game
}

func NewServer() *Server {
	server := new(Server)
	server.games = make(map[int]Game)
	return server
}

func (server *Server) Start() {
	fmt.Println("Ready to Dart !!")
	r := gin.Default()

	// les styles de jeu possibles
	r.GET("/styles", server.getStylesHandler) // retourne la liste des styles
	// creation du jeu (POST) -  fournit le type de jeu
	r.POST("/games", server.createNewGameHandler) // retourne un id
	// etat du jeu (GET)
	r.GET("/games/:gameId", server.findGameByIdHandler)
	// // creation du joueur (POST) -> retourne joueur
	r.POST("/games/:gameId/players", server.addPlayerToGameHandler)
	// // etat joueur
	// r.GET("/games/{gameId}/user/{userId}", server.userHandler).Methods("GET")
	//
	// // POST : etat de la flechette
	r.POST("/games/:gameId/darts", server.dartHandler)

	r.Run(":8080")
}

///GamesHandler
func (server *Server) createNewGameHandler(c *gin.Context) {
	var g common.GameRepresentation
	if c.BindJSON(&g) == nil {
		nextID := len(server.games) + 1

		theGame, err := gameFactory(g.Style)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "illegal content", "error": err.Error()})
			return
		}
		server.games[nextID] = theGame

		c.JSON(http.StatusOK, gin.H{"id": nextID, "game": theGame})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "illegal content"})
	}
}

func gameFactory(style string) (result Game, err error) {
	switch style {
	case common.GS_301.Code:
		result = NewGamex01(Optionx01{Score: 301, DoubleOut: false})
		return
	case common.GS_301_DO.Code:
		result = NewGamex01(Optionx01{Score: 301, DoubleOut: true})
		return
	case common.GS_501.Code:
		result = NewGamex01(Optionx01{Score: 501, DoubleOut: false})
		return
	case common.GS_501_DO.Code:
		result = NewGamex01(Optionx01{Score: 501, DoubleOut: true})
		return
	case common.GS_HIGH_3.Code:
		result = NewGameHighest(OptionHighest{Rounds: 3})
		return
	case common.GS_HIGH_5.Code:
		result = NewGameHighest(OptionHighest{Rounds: 5})
		return
	case common.GS_COUNTUP_300.Code:
		result = NewGameCountUp(OptionCountUp{Target: 300})
		return
	case common.GS_COUNTUP_500.Code:
		result = NewGameCountUp(OptionCountUp{Target: 500})
		return
	case common.GS_COUNTUP_900.Code:
		result = NewGameCountUp(OptionCountUp{Target: 900})
		return
	case common.GS_CRICKET.Code:
		result = NewGameCricket(OptionCricket{})
		return
	case common.GS_CRICKET_CUTTHROAT.Code:
		result = NewGameCricket(OptionCricket{CutThroat: true})
		return
	case common.GS_CRICKET_NOSCORE.Code:
		result = NewGameCricket(OptionCricket{NoScore: true})
		return
	default:
		err = errors.New("game of type " + style + " is not yet supported")
		return
	}
}

func (server *Server) findGameByIdHandler(c *gin.Context) {
	gameID, err := strconv.Atoi(c.Param("gameId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "illegal content", "error": err.Error()})
		return
	}
	log.WithFields(log.Fields{"gameID": gameID}).Info("flushing game w/ id")

	currentGame, ok := server.games[gameID]
	if !ok {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{"game": currentGame})
}

func (server *Server) addPlayerToGameHandler(c *gin.Context) {
	gameID, err := strconv.Atoi(c.Param("gameId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "illegal content", "error": err.Error()})
		return
	}

	log.Infof("flushing game w/ id {}", gameID)

	currentGame, ok := server.games[gameID]
	if !ok {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	var p common.PlayerRepresentation
	if c.BindJSON(&p) == nil {
		currentGame.AddPlayer(p.Name)
		c.JSON(http.StatusCreated, "http://localhost:8080/games/"+strconv.Itoa(gameID)+"/players")
	} else {
		c.JSON(http.StatusBadRequest, nil)
	}
}

func (server *Server) dartHandler(c *gin.Context) {
	gameID, err := strconv.Atoi(c.Param("gameId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "illegal content", "error": err.Error()})
		return
	}

	log.WithFields(log.Fields{"gameID": gameID}).Info("flushing game w/ id")

	currentGame, ok := server.games[gameID]
	if !ok {
		c.JSON(http.StatusNotFound, nil)
		return

	}

	var d common.DartRepresentation
	if c.BindJSON(&d) == nil {
		state, err := currentGame.HandleDart(common.Sector{Val: d.Sector, Pos: d.Multiplier})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, state)
		}

	} else {
		c.JSON(http.StatusBadRequest, nil)
	}
}

func (server *Server) getStylesHandler(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"styles": common.GS_STYLES})
}
