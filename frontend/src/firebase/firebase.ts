import { initializeApp } from "firebase/app"
import { getAuth } from "firebase/auth"

/*
For some magical reason, this does NOT work, even though this is exactly the same as
the firebaseConfig object below. Event checked using JSON.stringify(a) === JSON.stringify(b)
But it always shows incorrect api key.
const firebaseConfig = {
  apiKey: process.env.FIREBASE_API_KEY,
  authDomain: process.env.FIREBASE_AUTH_DOMAIN,
  projectId: process.env.FIREBASE_PROJECT_ID,
  storageBucket: process.env.FIREBASE_STORAGE_BUCKET,
  messagingSenderId: process.env.FIREBASE_MESSAGING_SENDER_ID,
  appId: process.env.FIREBASE_APP_ID,
}
*/

const firebaseConfig = {
  apiKey: "",
  authDomain: "",
  projectId: "",
  storageBucket: "",
  messagingSenderId: "",
  appId: "",
}

const app = initializeApp(firebaseConfig)

export const auth = getAuth(app)
export default app
