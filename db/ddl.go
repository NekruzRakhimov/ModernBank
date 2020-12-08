package db

const CreatingUsersTable = `Create table if not exists users(
	id integer primary key autoincrement,
	name text not null,
	surname text not null,
	age integer not null,
	gender text not null,
	role text not null,
	login text not null,
	password text not null,
	removed boolean not null default false
);`

const CreatingAccountsTable = `Create table if not exists accounts(
	id integer primary key autoincrement,
	user_id integer references users(id),
	amount integer not null,
	number text not null,
	system text not null,
	currency text not null,
	removed bool not null default false
);`

const CreatingATMsTable = `Create table if not exists atms(
	id integer primary key autoincrement,
	address text not null,
	status bool not null default true
);`

const CreatingTransactionsHistoryTable = `Create table if not exists transactionshistory(
	id integer primary key autoincrement,
	myaccount text not null,
	toaccount text not null,
	amount integer not null,
	data text not null,
	time text not null
);`
