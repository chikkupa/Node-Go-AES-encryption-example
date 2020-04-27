var encrypt = require("./node-encryption");

var text = "Hi, How are you?";
var key = "01234567890123450123456789012345";

const encrypedText = encrypt.encrypt(text, key);
console.log("Encrypted: ", encrypedText);
const decryptedText = encrypt.decrypt(encrypedText, key);
console.log("Decrypted: ", decryptedText);
