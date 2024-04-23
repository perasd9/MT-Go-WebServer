package server

import (
	"bytes"
	"fmt"
	handler "github/perasd9/MTWebServer/handlers"
	handlers "github/perasd9/MTWebServer/handlers/serverHandlers"
	database "github/perasd9/MTWebServer/infrastructure"
	"github/perasd9/MTWebServer/infrastructure/repos"
	"github/perasd9/MTWebServer/usecases"
	"io"
	"net"
	"os"
	"runtime"
	"strings"
	"sync"

	"github.com/joho/godotenv"
)

type MTServer struct {
	listener net.Listener
	router   *Router
}

// Constructing our Server
// Setting listener
func (s *MTServer) NewServer() *MTServer {
	err := godotenv.Load(".env")

	if err != nil {
		panic(err.Error())
	}
	ln, err := net.Listen("tcp", ":"+os.Getenv("ADDRESPORT"))

	if err != nil {
		panic(err.Error())
	}
	return &MTServer{
		listener: ln,
	}
}

// Starting server
func (mts *MTServer) Start() {
	//Closing listener in the end of method
	defer mts.listener.Close()

	//Limiting number of goroutines(lightweight threads) for defending thread problems
	runtime.GOMAXPROCS(10)
	//Synchronize goroutines into groups
	var wg sync.WaitGroup

	fmt.Printf("Server started on port %v \n", os.Getenv("ADDRESPORT"))
	fmt.Println(`
 ______     ______          ______     ______    __
/\  ___\   /\  __ \        /\  __ \   /\  == \  /\ \
\ \ \__ \  \ \ \_\ \       \ \  __ \  \ \   _/  \ \ \
 \ \__ __\  \ \_____\       \ \_\ \_\  \ \_\     \ \_\
  \/_____/   \/_____/        \/_/\/_/   \/_/      \/_/
  `)

	//Init router
	mts.router = NewRouter()

	//Defining database
	db := &database.MysqlDb{}

	db.GetDb()

	//Registering all repos
	programTypeRepository := repos.NewProgramTypeRepository(db)
	programRepository := repos.NewProgramRepository(db)
	exerciseRepository := repos.NewExerciseRepository(db)
	memberRepository := repos.NewMemberRepository(db)
	activityRepository := repos.NewActivityRepository(db)

	//Registering all usecases
	programTypeUsecase := usecases.NewProgramTypeUsecase(programTypeRepository)
	programUsecase := usecases.NewProgramUsecase(programRepository, activityRepository)
	exerciseUsecase := usecases.NewExerciseUsecase(exerciseRepository)
	memberUsecase := usecases.NewMemberUsecase(memberRepository)

	//Defining all controllers/handlers
	programTypeHandler := handler.NewProgramTypeHandler(programTypeUsecase)
	programHandler := handler.NewProgramHandler(programUsecase)
	exerciseHandler := handler.NewExerciseHandler(exerciseUsecase)
	memberHandler := handler.NewMemberHandler(memberUsecase)

	//Defining all possible routes by our API
	mts.router.AddRoute(*handlers.NewRoute("/programTypes", "GET"), programTypeHandler.GetAll)
	mts.router.AddRoute(*handlers.NewRoute("/programs", "DELETE"), programHandler.Delete)
	mts.router.AddRoute(*handlers.NewRoute("/programs", "POST"), programHandler.GetAll)
	mts.router.AddRoute(*handlers.NewRoute("/programs", "PUT"), programHandler.Update)
	mts.router.AddRoute(*handlers.NewRoute("/privateprograms", "POST"), programHandler.GetAllPrivatePrograms)
	mts.router.AddRoute(*handlers.NewRoute("/program", "POST"), programHandler.Add)
	mts.router.AddRoute(*handlers.NewRoute("/exercises", "GET"), exerciseHandler.GetAll)
	mts.router.AddRoute(*handlers.NewRoute("/login", "POST"), memberHandler.Login)
	mts.router.AddRoute(*handlers.NewRoute("/member", "POST"), memberHandler.Add)

	//Waiting for requests
	for {
		conn, err := mts.listener.Accept()

		if err != nil {
			fmt.Println("Error accepting connection ", err.Error())
			break
		}

		//Adding into synchronized group
		wg.Add(1)
		go func() {
			defer wg.Done()

			//Actual handling request
			handleConnection(conn, mts.router)
		}()
	}
	wg.Wait()
}

// Function for handling requests
func handleConnection(conn net.Conn, router *Router) {
	defer conn.Close()

	writer := io.Writer(conn)

	buffer := readBuffer(conn)
	buffer = bytes.Trim(buffer, "\x00")

	print("**********" + string(buffer) + "***********\n")

	if len(strings.TrimSpace(string(buffer))) == 0 {
		fmt.Println("Empty request received")

		return
	}

	//Creating requests by trimming data for http request
	request, er := handlers.NewRequest(string(buffer))

	if er {
		return
	}
	//Handling pure request by routing that request into supported handler
	responseJson, _ := router.Handle(request)

	//Write response back
	writer.Write([]byte(responseJson))
}

func readBuffer(conn net.Conn) []byte {
	//Init reader and writer for buffers

	//Init reader and writer for buffers
	buf := make([]byte, 8192)
	msg := make([]byte, 8192)

	for {
		line, err := conn.Read(buf)

		if err != nil {
			continue
		}
		msg = append(msg, buf[:line]...)

		if bytes.Contains(buf, []byte("}\x00\x00\x00")) {
			break
		}

		if bytes.Contains(buf, []byte("content-length")) {

			if bytes.Contains(buf, []byte("}")) {
				break
			}
			continue
		}

		if bytes.Contains(buf, []byte("\x00")) {

			break
		}
	}

	return msg
}
