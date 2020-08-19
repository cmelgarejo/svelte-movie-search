var nJwt = require("njwt");
module.exports = (req, res) => {
  try {
    console.log(
      nJwt.verify(
        req.query.jwt,
        "328c69c995a14a7f944623af20396c2c6f997ae806df4cf08eaf9f569cf8f8ad",
        "HS512"
      )
    );
  } catch (e) {
    console.log(e);
  }
  res.status(200).send(`Hello ${req.query.jwt}!`);
};
