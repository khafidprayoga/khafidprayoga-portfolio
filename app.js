const express = require("express")
require("dotenv").config()
const { AGE, LOCATION, EMAIL } = process.env


const app = express()
app.disable("x-powered-by")
app.get("/", (req, res) => {
    res.redirect("https://github.com/khafidprayoga")
})

app.get("/profile", (req, res) => {
    const response = {
        code: 200,
        success: true,
        message: "Hello World",
        data: {
            jobdesk: "Backend Developer",
            age: AGE,
            location: LOCATION,
            email: EMAIL,
        },
        stacks: [
            "Node.js",
            "GO",
            "AWS",
            "Docker",
        ],
    }
    res.status(response.code).send(response)
})

app.listen(3000, () => {
    console.log("Application running on port 3000")
})