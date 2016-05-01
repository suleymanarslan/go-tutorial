CREATE TABLE Interviews
(
    Id VARCHAR(50) PRIMARY KEY NOT NULL,
    IsFeatured BIT(1) DEFAULT b'0' NOT NULL,
    InterviewerId VARCHAR(50) NOT NULL,
    CategoryId VARCHAR(50) NOT NULL,
    Ranking INT(11) DEFAULT '-1' NOT NULL,
    Description VARCHAR(250) NOT NULL,
    Name VARCHAR(20) NOT NULL,
    IsActive BIT(1) DEFAULT b'1' NOT NULL,
    CONSTRAINT Interviews_InterviewCategories_Id_fk FOREIGN KEY (CategoryId) REFERENCES InterviewCategories (Id),
    CONSTRAINT Interviews_Interviewers_Id_fk FOREIGN KEY (InterviewerId) REFERENCES Interviewers (Id)
);
CREATE UNIQUE INDEX Interviews_Id_uindex ON Interviews (Id);
CREATE INDEX Interviews_InterviewCategories_Id_fk ON Interviews (CategoryId);
CREATE INDEX Interviews_Interviewers_Id_fk ON Interviews (InterviewerId);
CREATE TABLE Interviewers
(
    Id VARCHAR(50) PRIMARY KEY NOT NULL,
    UserId VARCHAR(50) NOT NULL,
    Ranking INT(11) NOT NULL,
    DateJoined DATETIME NOT NULL,
    CONSTRAINT Interviewers_Users_Id_fk FOREIGN KEY (UserId) REFERENCES Users (Id)
);
CREATE UNIQUE INDEX Interviewers_Id_uindex ON Interviewers (Id);
CREATE INDEX Interviewers_Users_Id_fk ON Interviewers (UserId);
CREATE TABLE InterviewSchedule
(
    Id VARCHAR(50) PRIMARY KEY NOT NULL,
    UserId VARCHAR(50) NOT NULL,
    InterviewerId VARCHAR(50) NOT NULL,
    InterviewDate DATETIME NOT NULL,
    CONSTRAINT InterviewSchedule_Interviewers_Id_fk FOREIGN KEY (InterviewerId) REFERENCES Interviewers (Id),
    CONSTRAINT InterviewSchedule_Users_Id_fk FOREIGN KEY (UserId) REFERENCES Users (Id)
);
CREATE UNIQUE INDEX InterviewSchedule_Id_uindex ON InterviewSchedule (Id);
CREATE INDEX InterviewSchedule_Interviewers_Id_fk ON InterviewSchedule (InterviewerId);
CREATE INDEX InterviewSchedule_Users_Id_fk ON InterviewSchedule (UserId);
CREATE TABLE Users
(
    Id VARCHAR(50) PRIMARY KEY NOT NULL,
    Username VARCHAR(50) NOT NULL,
    Password BINARY(1) NOT NULL,
    Email VARCHAR(50) NOT NULL,
    DateJoined DATETIME NOT NULL,
    IsActive BIT(1) NOT NULL,
    Summary VARCHAR(500)
);
CREATE TABLE InterviewCategories
(
    Id VARCHAR(50) PRIMARY KEY NOT NULL,
    CategoryName VARCHAR(50) NOT NULL,
    ParentId VARCHAR(50)
);
CREATE TABLE InterviewResult
(
    Id VARCHAR(50) NOT NULL,
    InterviewScheduleId VARCHAR(50) NOT NULL,
    Point INT(11) NOT NULL,
    Pros VARCHAR(200) NOT NULL,
    Cons VARCHAR(50) NOT NULL,
    EvaluationReport VARCHAR(250) NOT NULL,
    CONSTRAINT InterviewResult_InterviewSchedule_Id_fk FOREIGN KEY (InterviewScheduleId) REFERENCES InterviewSchedule (Id)
);
CREATE UNIQUE INDEX InterviewResult_Id_uindex ON InterviewResult (Id);
CREATE INDEX InterviewResult_InterviewSchedule_Id_fk ON InterviewResult (InterviewScheduleId);