package cli

import (
	"fmt"
	"wwchacalww/go-cem304/application"
	"wwchacalww/go-cem304/application/dto"
)

func Run(service application.UserServiceInterface, action string, userId string, userName string, userEmail string, userPassword string, userRole string) (string, error) {
	var result = ""

	switch action {
	case "create":
		var input dto.UserInput
		input.Name = userName
		input.Email = userEmail
		input.Password = userPassword
		input.Role = userRole

		user, err := service.Create(input)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("User ID %s with name %s and email %s has beenn created!", user.GetID(), user.GetName(), user.GetEmail())
		return result, err
	case "list":
		res, err := service.List()
		if err != nil {
			return result, err
		}

		for _, user := range res {
			text := fmt.Sprintf("Name: %s | email: %s | ID: %s | status: %t | role: %s \n", user.GetName(), user.GetEmail(), user.GetID(), user.GetStatus(), user.GetRole())
			result = result + text
		}
		return result, err
	default:
		res, err := service.FindByEmail(userEmail)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Name: %s, email: %s, password: %s, status: %t, role: %s", res.GetName(), res.GetEmail(), res.GetPassword(), res.GetStatus(), res.GetRole())
		return result, err
	}

}
