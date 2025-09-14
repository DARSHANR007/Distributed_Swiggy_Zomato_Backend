package main

import (
	"fmt"
	"food_delivery/DB/rdbms"
	"food_delivery/cmd/customer"
	"math/rand"
	"strings"
	"time"
)

func main() {
	// Initialize random seed
	rand.Seed(time.Now().UnixNano())

	// Connect to database
	db := rdbms.ConnectToDatabase()
	customerDB := rdbms.NewCustomerConnections(db)
	customer.InitCustomerDB(customerDB)

	// Lists of Indian names
	names := []string{
		"Arjun Sharma", "Priya Patel", "Vikram Singh", "Ananya Reddy", "Rahul Kumar",
		"Sneha Gupta", "Aditya Nair", "Kavya Iyer", "Rohan Mehta", "Deepika Joshi",
		"Karthik Rao", "Ritu Agarwal", "Siddharth Verma", "Pooja Malhotra", "Nikhil Shah",
		"Meera Krishnan", "Abhishek Mishra", "Shreya Bansal", "Varun Chopra", "Swati Pandey",
		"Gautam Pillai", "Neha Srinivasan", "Rajesh Chandra", "Divya Kulkarni", "Manish Tiwari",
		"Lakshmi Venkat", "Sandeep Bajaj", "Ritika Saxena", "Ashwin Desai", "Nandini Bhat",
		"Harish Gopal", "Shweta Khanna", "Ajay Ramesh", "Preeti Jain", "Sunil Naidu",
		"Madhuri Sethi", "Vishal Dixit", "Aarti Bhardwaj", "Naveen Gowda", "Sonali Arora",
		"Prakash Shetty", "Gayatri Menon", "Ravi Yadav", "Pallavi Hegde", "Sachin Kapoor",
		"Radha Sundaram", "Manoj Agrawal", "Shruti Dutta", "Krishna Murthy", "Sapna Ghosh",
		"Rakesh Trivedi", "Archana Vyas", "Deepak Raman", "Rashmi Singhal", "Vivek Bose",
		"Priyanka Das", "Sanjay Menon", "Nisha Rane", "Akash Jha", "Bhavana Kaul",
		"Amit Banerjee", "Tanvi Joshi", "Rohit Kulkarni", "Shilpa Nambiar", "Vishal Goyal",
		"Aditi Sinha", "Arjun Menon", "Riya Chatterjee", "Kunal Thakur", "Shreya Pandey",
		"Sameer Ahmed", "Pooja Raghavan", "Nitin Bhatt", "Prachi Agarwal", "Rajat Vashisht",
		"Kavitha Raman", "Saurabh Gupta", "Anita Sharma", "Varun Reddy", "Shrutika Iyer",
	}

	// Addresses from Bengaluru, Chennai, Mysore, Ooty
	baseAddresses := []string{
		// Bengaluru addresses
		"MG Road", "Brigade Road", "Koramangala", "Indiranagar", "Jayanagar",
		"Whitefield", "Electronic City", "BTM Layout", "HSR Layout", "Marathahalli",
		"Rajajinagar", "Malleswaram", "Basavanagudi", "Sadashivanagar", "JP Nagar",
		"Yelahanka", "Bannerghatta Road", "Sarjapur Road", "Hosur Road", "Old Airport Road",
		"RT Nagar", "Banaswadi", "Kalyan Nagar", "Hennur", "CV Raman Nagar",

		// Chennai addresses
		"T Nagar", "Anna Nagar", "Velachery", "Adyar", "Besant Nagar",
		"Mylapore", "Alwarpet", "Nungambakkam", "Egmore", "Guindy",
		"OMR Road", "ECR Road", "Porur", "Tambaram", "Chrompet",
		"Kilpauk", "Vadapalani", "Ashok Nagar", "Triplicane", "Royapettah",
		"Kodambakkam", "KK Nagar", "West Mambalam", "Saidapet", "Teynampet",

		// Mysore addresses
		"Saraswathipuram", "Kuvempunagar", "Vijayanagar", "Gokulam", "Hebbal",
		"Bogadi", "Nazarbad", "Chamundipuram", "Lakshmipuram", "Jayalakshmipuram",
		"V.V. Mohalla", "Ramakrishnanagar", "Dattagalli", "Yadavagiri", "Srirampura",

		// Ooty addresses
		"Charring Cross", "Commercial Road", "Garden Road", "Club Road", "Mysore Road",
		"Upper Bazaar", "Lower Bazaar", "Race Course Road", "Lake Road", "Elk Hill",
		"Coonoor Road", "Wellington Road", "Fernhill Road", "Observatory Road", "Church Road",
	}

	// Start timer (before registrations begin)
	start := time.Now()

	// Generate random data for each customer registration
	for i := 0; i < 1000; i++ {
		// Get random name
		randomName := names[rand.Intn(len(names))]

		// Generate random address with house number
		baseAddr := baseAddresses[rand.Intn(len(baseAddresses))]
		houseNumber := rand.Intn(999) + 1 // Random house number 1-999
		randomAddress := fmt.Sprintf("%d %s", houseNumber, baseAddr)

		// Generate random phone number (10 digits starting with 7,8,9)
		firstDigit := rand.Intn(3) + 7
		remainingDigits := rand.Intn(900000000) + 100000000
		randomPhone := fmt.Sprintf("%d%d", firstDigit, remainingDigits)

		// Generate email
		nameForEmail := strings.ToLower(strings.ReplaceAll(randomName, " ", "."))
		randomNumber := rand.Intn(9999) + 1
		randomEmail := fmt.Sprintf("%s%d@gmail.com", nameForEmail, randomNumber)

		// Keep the same password
		password := "securepassword"

		// Register the customer
		customer.Register(randomName, randomAddress, randomPhone, randomEmail, password)

		// Print progress every 100
		if (i+1)%100 == 0 {
			fmt.Printf("Registered %d customers\n", i+1)
		}
	}

	// Stop timer after all customers added
	elapsed := time.Since(start)

	fmt.Printf("Successfully registered 1000 customers with randomized data!\n")
	fmt.Printf("Time taken: %d ms\n", elapsed.Milliseconds())
}
