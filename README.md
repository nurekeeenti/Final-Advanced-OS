# Online Electronics Store

## Project Overview
This project is a final assignment for the **Advanced Databases (NoSQL)** course. It is a simple online electronics store where users can browse and purchase electronic products such as iPhones, laptops, and computers.

## Technologies Used
- **Backend:** Golang (Gin Framework)
- **Database:** MongoDB (NoSQL)
- **Frontend:** HTML, CSS, JavaScript
- **Authentication:** Email-based verification
- **Deployment:** Localhost (can be deployed to a cloud service)

## Features
- **User Authentication**: Users can register, log in, and verify their accounts via email.
- **Product Management**: View a list of available products, including their categories and brands.
- **Search Functionality**: Users can search for specific products.
- **Shopping Cart**: Add, update, and remove items from the cart.
- **Database Integration**: All data is stored in MongoDB.

## Installation
### Prerequisites
- Go installed on your machine.
- MongoDB database (local or cloud-based).
- SMTP credentials for email verification.

### Steps to Run the Project
1. Clone the repository:
   ```sh
   git clone https://github.com/your-repo-url.git
   cd your-repo-folder
   ```
2. Install dependencies:
   ```sh
   go mod tidy
   ```
3. Start the MongoDB database.
4. Update database connection details in `db.go`:
   ```go
   clientOptions := options.Client().ApplyURI("your-mongodb-connection-string")
   ```
5. Run the server:
   ```sh
   go run main.go
   ```
6. Open a browser and visit `http://localhost:8080`

## API Endpoints
### Authentication
- **POST /register** - Register a new user
- **POST /login** - Login user
- **POST /verify** - Verify email with a code

### Products
- **GET /api/products** - Get all products
- **GET /search?q=keyword** - Search for products

### Shopping Cart
- **GET /api/cart** - View cart
- **POST /api/cart/add** - Add product to cart
- **POST /api/cart/update** - Update product quantity
- **POST /api/cart/remove** - Remove product from cart

## Folder Structure
```
project-folder/
│-- main.go          # Main application file
│-- db.go            # Database connection setup
│-- auth.go          # Authentication routes
│-- cart.go          # Cart management routes
│-- products.go      # Product-related routes
│-- models/          # Data models
│-- static/          # HTML, CSS, JavaScript files
│-- templates/       # Frontend pages
```

## Future Improvements
- **User Roles:** Admin panel for product management.
- **Payment Integration:** Add payment gateways.
- **Improved UI:** Enhance the user experience with a modern design.

## Contributors
Nursultan Zhumagali
Aldiyar Kuandyk

## License
This project is open-source and available under the **MIT License**.

