package commands

import (
	"bufio"
	"fmt"
	"os"

	"github.com/alecthomas/kingpin"
	"golang.org/x/crypto/bcrypt"
)

type CostCommand struct {
}

func (command *CostCommand) run(context *kingpin.ParseContext) error {
	fmt.Print("Enter hashed password: ")
	reader := bufio.NewReader(os.Stdin)
	hashedPassword, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	cost, err := bcrypt.Cost([]byte(hashedPassword))
	if err != nil {
		return err
	}
	fmt.Println("Cost:", cost)
	return nil
}

func ConfigureCostCommand(app *kingpin.Application) {
	command := &CostCommand{}
	app.Command("cost", "Print the hashing cost used to create the given hash").Action(command.run)
}
