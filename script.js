document.getElementById('startBtn').addEventListener('click', () => {
    fetch('/start', { method: 'POST' })
        .then(response => response.json())
        .then(data => {
            document.getElementById('status').innerText = data.message;
        });
});

document.getElementById('stopBtn').addEventListener('click', () => {
    fetch('/stop', { method: 'POST' })
        .then(response => response.json())
        .then(data => {
            document.getElementById('status').innerText = `Tracking stopped. Social: ${data.SocialTime}, Web: ${data.WebTime}`;
        });
});

document.getElementById('reportBtn').addEventListener('click', () => {
    fetch('/sendReport', { method: 'POST' })
        .then(response => response.json())
        .then(data => {
            document.getElementById('status').innerText = data.message;
        });
});

