document.addEventListener('DOMContentLoaded', function () {
    const chatWindow = document.getElementById('chat-window');
    const messageInput = document.getElementById('message-input');
    const sendButton = document.getElementById('send-button');

    // 发送消息
    sendButton.addEventListener('click', function () {
        const message = messageInput.value.trim();
        if (message) {
            sendMessage(message);
            messageInput.value = '';
        }
    });

    // 获取消息
    setInterval(getMessages, 1000);

    // ...其他函数
});

// 发送消息函数
function sendMessage(message) {
    // ... Ajax发送消息到服务器
}

// 获取消息函数
function getMessages() {
    // ... Ajax获取新消息并显示
}