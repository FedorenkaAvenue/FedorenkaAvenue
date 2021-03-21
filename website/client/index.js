const dockerTestServerBoard = document.getElementById('test-docker');

const REQUEST_ERROR = 'request error :(';

window.addEventListener('load', getDockerStatus);

async function getDockerStatus() {
    try {
        const res = await fetch('http://localhost:1991/dockerps');
        const responce = await res.text();

        dockerTestServerBoard.innerHTML = responce;
    } catch(err) {
        dockerTestServerBoard.innerHTML = REQUEST_ERROR;
        console.error(err);
    }
}
