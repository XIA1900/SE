

/*
  i. owner: a owner
  ii. group_name
  iii. groups_avatar: group photo
  iv. description
  v. createdAt: create time
  vi. group id
*/
const date = new Date();


function basicInfo(groupName) {
    const result = {
        groupId: 7,
        owner: 'Lux',
        name: groupName,
        avatar: '/heroes/Lux_0.jpeg',
        description: 'Let\'s light it up',
        createdAt: date.getFullYear() + '-' + date.getMonth() + '-' + date.getDate(),
    };
    return result;
}

function getBasicInfo(req, res) {
    const params = req.query;
    const groupName = params.groupName;
    const result = basicInfo(groupName);
    return res.json({
        data: {
            list:result,
        },
    });
}

function analysis(groupName) {

}

function getAnalysis(req, res) {
    const params = req.query;
    const groupName = params.groupName;
    const result = analysis(groupName);
    return res.json({
        data: {
            list:result,
        },
    });
}

function member(groupName) {

}

function getMember(req, res) {
    const params = req.query;
    const groupName = params.groupName;
    const result = member(groupName);
    return res.json({
        data: {
            list:result,
        },
    });
}
 
function notification(groupName) {

}

function getNotification(req, res) {
    const params = req.query;
    const groupName = params.groupName;
    const result = notification(groupName);
    return res.json({
        data: {
            list:result,
        },
    });
}

function deleteGroup(req, res) {
    res.send({
        msg: 'Ok',
    });
}

function updateGroupInfo(req, res) {
    res.send({
        msg: 'Ok',
    });
}


export default {
    'GET /api/getBasicInfo': getBasicInfo,
    'GET /api/getAnalysis': getAnalysis,
    'GET /api/getMember': getMember,
    'GET /api/getNotification': getNotification,
    'POST /api/updateGroupInfo': async (req, res) => {
        res.send({
          message: 'Ok',
        });
      },
    'POST /api/deletGroup': async (req, res) => {
        res.send({
          message: 'Ok',
        });
      },
  };
  