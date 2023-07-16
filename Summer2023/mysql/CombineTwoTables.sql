SELECT FirstName, LastName, City, State 
FROM Person left join Address 
ON Person.PersonId = Address.PersonId;