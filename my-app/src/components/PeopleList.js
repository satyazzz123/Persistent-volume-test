import React, { useState, useEffect } from 'react';
import axios from 'axios';

const PeopleList = () => {
  const [people, setPeople] = useState([]);

  useEffect(() => {
    const fetchPeople = async () => {
      try {
        const response = await axios.get(`${process.env.REACT_APP_API_URL}`);
        setPeople(response.data);
      } catch (error) {
        console.error('Error fetching people:', error);
      }
    };

    fetchPeople();
  }, []);

  return (
    <div>
      <h2>People List</h2>
      <ul>
        {people.map((person) => (
          <li key={person._id}>
            {person.name} - {person.age} years old
          </li>
        ))}
      </ul>
    </div>
  );
};

export default PeopleList;
