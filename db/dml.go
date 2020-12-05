package db

const UpadateAccountsAmount = `UPDATE accounts set amount = $1 where number = $2`

const SelectAccountByUser_id = `SELECT * FROM accounts where user_id = ($1)`

const SelectAccountByNumber = `SELECT * FROM accounts WHERE number = ($1)`

const AddTransactionHistory = `INSERT INTO transactionshistory(myaccount, toaccount, amount, data, time) 
VALUES(($1), ($2), ($3), ($4), ($5))`

const AddUser = `INSERT INTO users(name, surname, age, gender, role, login, password) 
VALUES(($1), ($2), ($3), ($4), ($5), ($6), ($7))`

const AddAccount = `INSERT INTO accounts(user_id, amount, number, system, currency) 
VALUES(($1), ($2), ($3), ($4), ($5))`

const UpadateUser = `UPDATE users set name = $1, surname = $2, age = $3, gender = $4, role = $5, login = $6, password = $7 where id = $8`

const UpadateAccount = `UPDATE accounts set id = $1, user_id = $2, system = $3, currency = $4 where id = $5`

const UpadateATM = `UPDATE atms set address = $1 where id = $2`


const RemoveAccount = `UPDATE accounts set removed = true where id = $1`

const RemoveATM = `UPDATE atms set status = false where id = $1`


const RemoveUser = `UPDATE users set removed = true where id = $1`

const SelectTransactionsHistory = `SELECT * FROM transactionshistory WHERE myaccount = ($1)`