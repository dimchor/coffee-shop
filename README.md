# Coffee Shop
This is a first attempt in creating a simple CRUD app. This is an assignment for 
university course called "Special Topics in Software Engineering".

A user may:
- register to the platform,
- login to the platform,
- change the password,
- add coffee products to shopping cart and
- order coffee products.

The admin is able to add products to the platform using a special interface.

The project was developed using the SCUM methodology and it was tracked using 
GitHub Projects.

![Capture](https://github.com/user-attachments/assets/147a2bcc-db5c-4896-b355-dff22510682e)

# Setup
Use `docker-compose build` and `docker-compose up -d` to run the app.

A `.env` file is used to store the database's password. Create the file in this 
directory and assign the password to the `DB_ROOT_PASSWORD` environmental 
variable.
