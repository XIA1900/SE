export default {
    'POST  /api/createGroup': (req, res) => {
      console.log("params");
      console.log(req.query);
      res.send({
        data: {
          message: 'Ok',
        },
      });
    },
    'POST  /api/createPost': (req, res) => {
        res.send({
          data: {
            message: 'Ok',
          },
        });
      },
  };

  