package cmd

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

var generateMealCmd = &cobra.Command{
	Use:   "meal",
	Short: "Generate a random meal plan",
	Long:  `Generates a random meal suggestion for the week.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		rand.Seed(time.Now().UnixNano())
		fmt.Println("Your Meals:")

		for _, options := range meals {
			fmt.Printf("%s\n", options[rand.Intn(len(options))])
		}
	},
}

var meals = map[string][]string{
	"Quick & Easy": {
		"Sweet & Sticky Five-Spice Chicken Burger",
		"Shredded Hoisin Chicken Wraps",
		"Herby Crumbled Chicken (Panko Chicken)",
		"Chicken 'Meteorite' Nugget and Mash",
		"Spicy Honey Glazed Chicken",
		"King’s Coronation Chicken Sandwich",
		"Bash-Mi Style Chicken Sandwich",
	},
	"Burger, Wraps, Tacos, or Pie": {
		"Chicken Curry Pasty with Chips and Mango-Sriracha Dip",
		"Truffled Chicken, Bacon & Mushroom Pie",
		"Bulgogi Pork and Pepper Tacos",
		"Mango Chutney Glazed Chicken Wraps",
		"Hawaiian Inspired Bacon Cheeseburger",
		"Crispy Buffalo Chicken Tacos and Chips",
		"The Truffle Cheese",
		"Korean BBQ Style Burger and Sticky Chicken",
		"Ghoulish Chicken Goujons",
		"Crispy Chicken Strip Tacos",
		"Ultimate Bacon Cheeseburger and Chips",
		"Pig Under Blanket Sausage and Bacon Burger",
		"Crispy Chicken and Gravy Burger",
		"The Coronation Chick-King",
		"Fried Chicken Burger and Cheesy Wedges",
	},
	"Noodles, Rice or Pasta": {
		"Carbonara",
		"Thai Style Pork Rice Bowl",
		"Sweet Sticky Chicken",
		"Pesto Chicken Traybake and Garlic Rice",
		"Sticky Baked Hoisin Chicken Thighs",
		"Soy, Ginger & Lime Pork Meatballs",
		"Speedy Chicken Noodles",
		"Sweet Chilli Beef",
		"Bulgogi Chicken Stir-Fry",
		"Ultimate Chicken Tikka and Cumin Rice",
		"Honey Garlic Pork Meatballs",
		"Quick Butter Chicken Masala",
		"Teriyaki Lemongrass Beef",
		"Quick & Sticky Beef",
		"Teriyaki Sesame Chicken",
		"Teriyaki Ginger Beef Noodles",
		"Stir-Fried Hoisin Chicken Noodles",
	},
	"Other Stuff": {
		"Roasted Chicken with Chili & Chive Sauce",
		"Crispy Japanese Style Fried Chicken",
		"Pan-Fried Garlic Butter Chicken Thighs",
		"Pork Meatballs in Creamy Chive Sauce",
		"Rosemary Roast Chicken & Tomato 'Asteroids'",
		"Caribbean Style Jerk Chicken Thighs",
		"Korma Style Spiced Chicken and Pepper Skewers",
		"Glazed Peri Peri Chicken Thighs",
		"Peri Peri Chicken Breast Traybake",
		"Tandoori Chicken and Potato Traybake",
		"Herby Parsley Chicken",
		"Harissa Chicken and Vegetable Traybake",
		"Firecracker Style Chicken",
		"Fajita Spiced Chicken",
		"Pork Steaks with Mustard Chive Sauce",
		"Pork Steak and Creamy Peppercorn Sauce",
		"Beef Meatballs in Onion & Redcurrant Sauce",
		"Honey Mustard Sausage and Red Onion Gravy",
		"Spooky Bacon and Cumberland Sausage Mummies",
	},
}
