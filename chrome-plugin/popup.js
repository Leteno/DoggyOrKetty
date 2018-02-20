var backgroundPage = chrome.extension.getBackgroundPage();

var setStat = function(jsonData) {

    var json = JSON.parse(jsonData);
    var hungry = document.getElementById('hungry');
    var thirsty = document.getElementById('thirsty');
    var mood = document.getElementById('mood');
    hungry.innerHTML = json.hungry;
    thirsty.innerHTML = json.thirsty;
    mood.innerHTML = json.mood;
};

document.addEventListener('DOMContentLoaded', function() {
	backgroundPage.localApi('stat', setStat);
    });
