
// ##################
// Notification part
// ##################
Notification.requestPermission();
chrome.notifications.onClicked.addListener(function(notificationId) {
        chrome.tabs.create({url: notificationId});
        chrome.notifications.clear(notificationId);
	localApi('feed?hungry=20');
    });
var showNotification = function(title, message, url) {
    console.log('show notification message: ' + message + ', url: ' + url);
    var option = {
        type: 'basic',
        title: title,
        message: message,
        iconUrl: 'Liu.png',
    };
    var notificationId = url==undefined ? 'https://www.zhihu.com' : url;
    chrome.notifications.clear(notificationId);
    chrome.notifications.create(notificationId, option);
}

// ##################
// api part
// ##################

var localApi = function(dest, callback) {
    var uri = '127.0.0.1'
    var port = 8000
    var xhr = new XMLHttpRequest();
    xhr.onload = function() {
        callback(xhr.response);
    };
    xhr.open('GET', 'http://127.0.0.1:8000/' + dest);
    xhr.send();
};

// ###########################
// #  new logic start at here
// ###########################

var jobParser = function(json) {
    what = json.what;
    console.debug("jobParser: " + what);
    switch(what) {
    case "nothing":
	break;
    case "notification":
	title = json.content.title;
	body = json.content.body;
	target_url = json.content.target_url;
	showNotification(title, body, target_url);
	break;
    }
};

var imQuery = function() {
    localApi('im', function(data) {
	    var json = JSON.parse(data);
	    jobParser(json);
	});
};
ten_seconds = 10 * 1000;
setInterval(imQuery, ten_seconds);
imQuery();