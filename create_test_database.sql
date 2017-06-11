CREATE TABLE tickets(
	id INTEGER PRIMARY KEY,
	owner INTEGER,
	status TEXT,
	priority INTEGER,
	headline TEXT,
	FOREIGN KEY(owner) REFERENCES users(id)
);

CREATE TABLE comments(
	id INTEGER PRIMARY KEY,
	ticket INTEGER,
	created DATETIME DEFAULT CURRENT_TIMESTAMP,
	author INTEGER,
	public INTEGER,
	pinned INTEGER,
	highlight INTEGER,
	body TEXT,
	FOREIGN KEY(ticket) REFERENCES tickets(id)
	FOREIGN KEY(author) REFERENCES users(id)
);

CREATE TABLE tags(
	ticket INTEGER,
	tag TEXT,
	FOREIGN KEY(ticket) REFERENCES tickets(id)
);

CREATE TABLE users(
	id INTEGER PRIMARY KEY,
	email TEXT,
	name TEXT,
	enabled INTEGER,
	password TEXT,
	salt TEXT
);

INSERT INTO users VALUES(1, 'paul@example.com', 'Paul', 1, '$2a$10$MBxQWcgk5U7JtABPz8imQ.u0J79z2JSoR8dvpuS4VY0n8q6Xyz8Y6', 'abcd1234');
INSERT INTO users VALUES(2, 'john@example.com', 'John', 1, '$2a$10$vJbH7BXjMhtLRbADn4F2Ue8sXEq2Rfj4tERHuRJoN7K86CNrGbkmW', '9876wxyz');

INSERT INTO tickets VALUES(1, 1, 'Open', 3, 'Joe in Accounting can''t print');
INSERT INTO comments VALUES(NULL, 1, '2017-06-09 10:30:58', 1, 0, 0, 0, 'Joe left a voicemail. He hasn''t been able to print since this morning. #print');
INSERT INTO comments VALUES(NULL, 1, '2017-06-09 10:38:14', 2, 0, 0, 0, 'I emptied his print queue, and restarted the spooler. I emailed him to try printing again.');

INSERT INTO tickets VALUES(2, 1, 'Open', 3, 'Packet loss at Birmingham office');
INSERT INTO comments VALUES(NULL, 2, '2017-06-10 11:30:28', 1, 0, 0, 0, 'Icinga alerted about packet loss at Birmingham. #birmingham #network');
INSERT INTO comments VALUES(NULL, 2, '2017-06-10 11:38:44', 1, 0, 0, 0, 'They''re dropping about 3% of packets. MTR points to the provider''s device upstream of our router.');

Insert INTO tags VALUES(1, "print");
Insert INTO tags VALUES(2, "birmingham");
Insert INTO tags VALUES(2, "network");
