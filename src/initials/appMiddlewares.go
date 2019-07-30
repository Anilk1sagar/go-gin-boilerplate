package initials

import (
	"fmt"
)

// InitializeMiddlewares for App
func (a *App) InitializeMiddlewares() {

	fmt.Println("**==== Initializing Middlewares ====**")

	// /* ========================== Initialize Firebase Admin ======================= */
	// ctx := context.Background()
	// conf := &firebase.Config{
	// 	DatabaseURL: "https://gobaseapp-82fd5.firebaseio.com",
	// }

	// // Fetch the service account key JSON file contents
	// opt := option.WithCredentialsFile("gobaseapp-82fd5-firebase-adminsdk.json")

	// // Initialize the app with a service account, granting admin privileges
	// admin, err := firebase.NewApp(ctx, conf, opt)

	// if err != nil {
	// 	utils.Logger().Fatalln("Error initializing Firebase app:", err.Error())
	// }

	// // Initialize Firebase db
	// firebaseDb, err := admin.Database(ctx)

	// if err != nil {
	// 	utils.Logger().Fatalln("Error initializing Firebase database:", err)
	// }

	// // Assigning firebasedb to app
	// a.FirebaseDb = firebaseDb

	// fmt.Println("** Firebase Initialize: ** ", admin, firebaseDb)

}
