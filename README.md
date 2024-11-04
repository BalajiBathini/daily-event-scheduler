
---

# Daily Event Scheduler

## Overview

The Daily Event Scheduler is a simple scheduling application that manages events within a 24-hour day. The application allows users to schedule events by specifying their start and end times while ensuring that no events overlap. It features a user-friendly interface for adding and viewing scheduled events.

## Table of Contents
- [Features](#features)
- [Setup Instructions](#setup-instructions)
- [Core Components](#core-components)
- [Event Management](#event-management)
- [User Interface](#user-interface)
- [Testing the Application](#testing-the-application)
- [Error Handling](#error-handling)
- [Thought Process](#thought-process)
- [Contributing](#contributing)
- [License](#license)

## Features
- Schedule events within a 24-hour format (0-23).
- Prevent overlapping events.
- Display all scheduled events in a clear manner.
- Input validation to ensure valid event times.

## Setup Instructions

1. **Clone the Repository**:
   ```bash
   git clone <repository-url>
   cd daily-event-scheduler
   ```

2. **Install Dependencies**:
   If you're using a frontend framework (like React, Angular, etc.), make sure to install any necessary dependencies. For a simple HTML/CSS/JavaScript project, no additional setup is required.

3. **Open the Application**:
   - Open `index.html` in your web browser to view the application.

## Core Components

### Scheduler Class
The `Scheduler` class is the backbone of the application. It manages the events and includes the following methods:

- **addEvent({start_time: number, end_time: number})**:
  - Checks for overlapping events before adding a new event.
  - Returns `true` if the event is added successfully; otherwise, returns `false`.

- **getEvents()**:
  - Returns an array of all scheduled events.

### Event Representation
Events are represented by their start and end times, with whole numbers between 0 and 23 denoting hours in a day.

## Event Management

The `addEvent` method will:
- Validate the start and end times.
- Ensure no overlapping with existing events.
- Add the event to the `events` array if valid.

The `getEvents` method returns the list of all scheduled events.

## User Interface

The user interface allows users to:
- Input start and end times for events.
- View all scheduled events displayed in a clear list or timeline format.
- Receive feedback if an event overlaps or is invalid.

### Example Scenarios
- If events (2, 5) and (7, 9) are scheduled:
  - Adding (1, 3) fails (overlaps with 2-5).
  - Adding (5, 7) succeeds (fits between events).
  - Adding (8, 10) fails (overlaps with 7-9).

## Testing the Application

- Open the application in a web browser.
- Test the event scheduling by entering various start and end times.
- Ensure the application correctly handles overlaps and displays events.

## Error Handling
- The application includes input validation to prevent invalid time entries.
- Users will receive alerts for overlapping events or incorrect input formats.

## Thought Process
In developing this application, I focused on creating a straightforward and efficient way to manage daily events. The main challenge was ensuring no overlaps occurred when adding events. I implemented a simple algorithm to check for overlaps, which allows for real-time validation before scheduling an event. The user interface is designed to be intuitive, enabling easy interaction for adding and viewing events.

## Contributing
Contributions are welcome! If you have suggestions or improvements, feel free to fork the repository and submit a pull request.

## License
This project is licensed under the MIT License. See the LICENSE file for details.

---
