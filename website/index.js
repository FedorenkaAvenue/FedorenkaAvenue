const dockerTestServerBoard = document.getElementById('test-docker');

const REQUEST_ERROR = 'request error :(';

window.addEventListener('load', getDockerStatus);

async function getDockerStatus() {
    try {
        const res = await fetch('https://fedorenka.online:1991/docker/ps');
        const responce = await res.text();

        if (res.status === 406) return dockerTestServerBoard.innerHTML = 'docker error :(';

        dockerTestServerBoard.innerHTML = responce;
    } catch(err) {
        dockerTestServerBoard.innerHTML = REQUEST_ERROR;
    }
}
