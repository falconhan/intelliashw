package main

import (
	"errors"
	"fmt"
)

type animals interface {
	amountFeedGetter
	animalNameGetter
	animalSkinGetter
}

type amountFeedGetter interface {
	getAmountFood() (float64, error)
}

type animalNameGetter interface {
	getAnimalName() string
}

type animalSkinGetter interface {
	getAnimalSkin() string
}

type animal struct {
	voice              string
	suitabilityForFood bool
	skin               string
	name               string
	weight             float64
	weightForFood      float64
	weightFood         float64
}

func (a animal) getAnimalSkin() string {
	return a.skin
}

func (a animal) getAnimalName() string {
	return a.name

}
func (a animal) getAmountFood() (float64, error) {

	if err := validateType(a); err != nil {
		return 0, fmt.Errorf("wrong skin")
	}

	if err := validateWeight(a); err != nil {
		return 0, fmt.Errorf("failed to calculate: %w", err)
	}

	if err := validateSuitabilityForFood(a); err != nil {
		return 0, fmt.Errorf("wrong suitability for food: %w", err)
	}

	foodForOneKg := a.weightFood / a.weightForFood
	ourFood := foodForOneKg * a.weight
	return ourFood, nil
}

type dog struct {
	animal
}

type cat struct {
	animal
}

type cow struct {
	animal
}

func validateType(a animal) error {
	if (a.voice != "Wow" && a.skin == "dog") || (a.voice != "May" && a.skin == "cat") || (a.voice != "Myyy" && a.skin == "cow") {
		return errors.New("bad type")
	}
	return nil
}

var smallWeightErr = errors.New("bad weight")

func validateWeight(a animal) error {
	if a.weight <= 0 {
		return smallWeightErr
	}
	return nil
}

var suitabilityForFoodErr = errors.New("false suitability for food")

func validateSuitabilityForFood(a animal) error {
	if (a.suitabilityForFood == true && (a.skin == "dog" || a.skin == "cat")) || (a.suitabilityForFood == false && a.skin == "cow") {
		return suitabilityForFoodErr
	}
	return nil
}

func main() {

	animalsOnTheFarm := []animals{
		dog{animal{voice: "Wow", suitabilityForFood: true, skin: "dog", name: "Sharik", weight: 5, weightFood: 10, weightForFood: 15}},
		dog{animal{voice: "Wow", suitabilityForFood: false, skin: "dog", name: "Mukhtar", weight: 8, weightFood: 19, weightForFood: 5}},
		cat{animal{voice: "May", suitabilityForFood: false, skin: "cat", name: "Barsic", weight: 2, weightFood: 7, weightForFood: 3}},
		cow{animal{"Myyy", true, "cow", "Marta", 200, 12, 31}},
	}

	for _, a := range animalsOnTheFarm {
		infoAnimal, err := a.getAmountFood()

		if errors.Is(err, suitabilityForFoodErr) {
			fmt.Println(a.getAnimalSkin(), a.getAnimalName(), ":", "does not correct suitability for food :", err)
			continue
		}

		if errors.Is(err, smallWeightErr) {
			fmt.Println(a.getAnimalSkin(), a.getAnimalName(), ":", "small weight :", err)
			continue
		}

		if err != nil {
			fmt.Println("for", a.getAnimalSkin(), a.getAnimalName(), ": bad skin")
			return
		}

		fmt.Println("Amount of feed for", a.getAnimalSkin(), a.getAnimalName(), any(infoAnimal))
	}

}
