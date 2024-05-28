import React from 'react';
import './App.css';
import PeopleList from './components/PeopleList';
import AddPersonForm from './components/AddPersonForm';

const App = () => {
  return (
    <div className="App">
      <h1>People Manager</h1>
      <AddPersonForm />
      <PeopleList />
    </div>
  );
};

export default App;
