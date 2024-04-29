package main

import (
    "errors"
    "log"
    "fmt"
    "github.com/charmbracelet/huh"
)

func main() { 
    var (
        burger string
        toppings []string
        sauceLevel int
        name string
        instructions string
        discount bool
    )

    form := huh.NewForm(
        huh.NewGroup(

        huh.NewSelect[string]().
            Title("Choose your borgar").
            Options(
                huh.NewOption("Charmburger Classic", "clasic"),
                huh.NewOption("Chick-fil-A", "chickwich"),
                huh.NewOption("Fishburger", "fishburger"),
                huh.NewOption("Charmpossible", "charmpossible"),
            ).
            Value(&burger),
        huh.NewMultiSelect[string]().
            Title("Toppings").
            Options(
                huh.NewOption("Lettuce", "lettuce").Selected(true),
                huh.NewOption("Tomatoes", "tomatoes").Selected(true),
                huh.NewOption("Jalapeños", "jalapeños").Selected(true),
                huh.NewOption("Cheese", "cheese"),
                huh.NewOption("Vegan Cheese", "vegan cheese"),
                huh.NewOption("Nutella","nutella"),
            ).
            Limit(4).
            Value(&toppings),
        huh.NewSelect[int]().
        Title("How much Charm Sauce do you want?").
        Options(
            huh.NewOption("None", 0),
            huh.NewOption("A little", 1),
            huh.NewOption("A lot", 2),
        ).
        Value(&sauceLevel),
        ),
        huh.NewGroup(
            huh.NewInput().
                Title("What's your name?").
                Value(&name).
                Validate(func(str string) error {
                    if str == "Hawkwhisper" {
                        return errors.New("Sorry, we don't serve customers named Hawkwhisper.")
                    }

                    return nil
                }),

                huh.NewText().
                    Title("Special Instructions").
                    CharLimit(500).
                    Value(&instructions),
                huh.NewConfirm().
                    Title("Would you like 15% off?").
                    Value(&discount),
        ),
    )

    err := form.Run()
    if err != nil {
        log.Fatal(err)
    }

    if !discount {
        fmt.Println("What? You didn't take the free discount!?")
    }

    receipt := `
        Your receipt
    --------------------
    Burger: %s
    toppings: %s
    Sauce level: %d
    Customer name: %s
    Special Instructions: %s
    Discount: %v
    --------------------
        Thank you~!
        Come again!
    `
    formattedReceipt := fmt.Sprintf(receipt, burger, toppings, sauceLevel, name, instructions, discount)
    fmt.Println(formattedReceipt)
}
