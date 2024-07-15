document.addEventListener("DOMContentLoaded", function() {
    fetch("/api/schedule")
        .then(response => response.json())
        .then(data => {
            const scheduleDiv = document.getElementById("schedule");
            // Render schedule data into the scheduleDiv
            // Implement your rendering logic here
        });
});
