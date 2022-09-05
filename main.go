package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Conferencia de GO";
	const conferenceTicket = 50;
	var remainingTickets uint = 50;
	greetUsers(conferenceName);
	fmt.Printf("Acá podrás obtener tus Tickets, actualmente quedan %v disponibles de %v \n", remainingTickets, conferenceTicket);

	var bookings []string;
	
	for {
		var userName string;
		var userLastName string;
		var email string;
		var userTickets uint;

		fmt.Println("Ingresa tu nombre");
		fmt.Scan(&userName)
	
		fmt.Println("Ingresa tu apellido");
		fmt.Scan(&userLastName)
	
		fmt.Println("Ingresa tu email");
		fmt.Scan(&email)
	
		fmt.Println("Ingresa cuantos tickets quieres");
		fmt.Scan(&userTickets)

		isValidName := len(userName) >= 2 && len(userLastName) >= 2;
		isValidEmail := strings.Contains(email, "@");
		isValidTickets := userTickets > 0 && userTickets <= remainingTickets; 
		
		if isValidName && isValidEmail && isValidTickets {
			remainingTickets = remainingTickets - userTickets;
			bookings = append(bookings, userName + " " + userLastName);
			fmt.Printf("Gracias %v %v por adquirir %v tickets, hemos enviado el comprobante a %v, ten un gran día \n", userName, userLastName ,userTickets, email);
			fmt.Printf("Quedan %v tickets disponibles de %v \n", remainingTickets, conferenceTicket);
			firstNames := []string{};
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
				} 
				fmt.Printf("Estas son las personas que han comprado tickets %v \n", firstNames);
				if remainingTickets == 0 {
					fmt.Println("Ya no quedan tickets disponibles");
					break;
				}
		} else {
			if !isValidName {
				fmt.Println("Ingresa un nombre y apellido válido, el que ingresaste es muy corto");
			}
			if !isValidEmail {
				fmt.Println("Ingresa un email válido, el que ingresaste no tiene @");
			}
			if !isValidTickets {
				fmt.Println("Ingresa un número de tickets válido, el que ingresaste es mayor al disponible o es 0");
			}
		}
	}
}

func greetUsers(conferenceName string) {
	fmt.Printf("Bienvenido a nuestra aplicación para ingresar a %v, lo mejor que se nos ocurrió por ahora\n", conferenceName);
}