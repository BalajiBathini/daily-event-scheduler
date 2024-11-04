const form = document.getElementById('event-form');
const eventsList = document.getElementById('events');

// Function to fetch events from the backend
function fetchEvents() {
    fetch('http://localhost:8080/events')
        .then(response => response.json())
        .then(data => {
            eventsList.innerHTML = '';
            data.forEach(event => {
                const listItem = document.createElement('li');
                listItem.textContent = `Event: ${event.start_time} - ${event.end_time}`;
                eventsList.appendChild(listItem);
            });
        })
        .catch(error => console.error('Error fetching events:', error));
}

// Handle form submission
form.addEventListener('submit', (e) => {
    e.preventDefault(); // Prevent default form submission

    const startTime = parseInt(form.start_time.value);
    const endTime = parseInt(form.end_time.value);

    // Create event object
    const eventData = {
        start_time: startTime,
        end_time: endTime,
    };

    // Send POST request to add event
    fetch('http://localhost:8080/events', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(eventData),
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        return response.json();
    })
    .then(data => {
        console.log('Event added:', data);
        fetchEvents(); // Refresh the event list
        form.reset(); // Reset form fields
    })
    .catch(error => {
        console.error('Error adding event:', error);
        alert('Error adding event. Please check the times and try again.');
    });
});

// Fetch events on page load
fetchEvents();
