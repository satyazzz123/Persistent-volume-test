const express = require('express');
const mongoose = require('mongoose');

const app = express();
app.use(express.json());

// Connect to MongoDB
mongoose.connect('mongodb://0.0.0.0:27017', {
  useNewUrlParser: true,
  useUnifiedTopology: true,
});

// Define a schema for the data
const personSchema = new mongoose.Schema({
  name: String,
  age: Number,
});

// Create a model based on the schema
const Person = mongoose.model('Person', personSchema);

// GET API endpoint to retrieve all people
app.get('/people', async (req, res) => {
  try {
    const people = await Person.find();
    res.json(people);
  } catch (err) {
    res.status(500).json({ message: err.message });
  }
});

// POST API endpoint to create a new person
app.post('/people', async (req, res) => {
  const person = new Person({
    name: req.body.name,
    age: req.body.age,
  });

  try {
    const newPerson = await person.save();
    res.status(201).json(newPerson);
  } catch (err) {
    res.status(400).json({ message: err.message });
  }
});

app.listen(8000, () => {
  console.log('Server started on port 6000');
});