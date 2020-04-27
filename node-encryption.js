var crypto = require("crypto");
const iv = "1100273601254817";
// Decrypts a string coming from the go encrytion method
exports.decrypt = function (content, key) {
  // Create the cipher for decrypting the passed in string. We use the IV that was passed in and would have been encoded to hex.
  // We therefore need get the IV back into the string used by the encryption
  var decipher = crypto.createDecipheriv(
    "aes-256-cbc",
    key.toString("binary"),
    iv
  );

  // Decrypt the string passed in
  var decrypted = decipher.update(content, "hex", "utf8");
  decrypted += decipher.final("utf8");

  // return the now decrypted string
  return decrypted;
};

// Encypts a string into an AES format compatible with node
exports.encrypt = function (content, key) {
  // Create the cipher for the encyption
  cipher = crypto.createCipheriv("aes-256-cbc", key.toString("binary"), iv);

  // Set the encoding to hex. Some people choose base64 but I prefer hex as it will generate numbers which are more predictable
  cipher.setEncoding("hex");
  cipher.write(content);
  cipher.end();

  cipher_text = cipher.read();

  // Create the encrypted text string using the "$" as a seperator of the IV which is needed for the decryption
  return cipher_text; // + "$" + iv.toString('hex');
};
