import { request } from 'umi';

/*
search something: 

input: values is a string

return: most related 10 articles, same properties as /api/queryList
*/

export async function checkMember(values) {
  return request('/api/checkMember', {
    params: values,
    method: 'POST',
  });
}


export async function quitGroup(values) {
  return request('/api/quitGroup', {
    params: values,
    method: 'POST',
  });
}

export async function joinGroup(values) {
  return request('/api/joinGroup', {
    params: values,
    method: 'POST',
  });
}

export async function queryCurrent(params) {
  return request('/api/user/getuserinfo?username='+params.username, {
  //return request('/api/currentUserDetail', {  
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
    credentials: 'include',
  });
}

export async function userUpdate(body) {
  return request('/api/user/update', {
  //return request('/api/currentUserDetail', {  
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    credentials: 'include',
    data: body,
  });
}

export async function getPersonalCollection(params) {
  //return request('/api/getPersonalCollection');
  return request('/api/articlefavorite/get?pageno='+params.pageNO+'&pagesize='+params.pageSize, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
    credentials: 'include',
  });
}




export async function getPersonalFollower(values) {
  //return request('/api/getPersnalFollower');
  return request('/api/user/followers', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    credentials: 'include',
  });
}

export async function removeFollower(values) {
  return request('/api/removeFollower', {
    params: values,
    method: 'POST',
  });
}

export async function getPersonalFollowing(values) {
  //return request('/api/getPersonalFollowing');
  return request('/api/user/followees', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    credentials: 'include',
  });
}
  
export async function removeFollowing(values) {
  return request('/api/removeFollowing', {
    params: values,
    method: 'POST',
  });
}

export async function addFollowing(values) {
  return request('/api/addFollowing', {
    params: values,
    method: 'POST',
  });
}

export async function getPersonalBlacklist(values) {
  return request('/api/getPersonalBlacklist');
}

export async function removeBlacklist(values) {
  return request('/api/removeBlacklist', {
    params: values,
    method: 'POST',
  });
}

export async function addBlacklist(values) {
  return request('/api/addBlacklist', {
    params: values,
    method: 'POST',
  });
}

export async function changePassword(values) {
  return request('/api/changePassword', {
    params: values,
    method: 'POST',
  });
}

export async function removeLike(values) {
  return request('/api/removeLike', {
    params: values,
    method: 'POST',
  });
}

export async function createLike(params) {
  return request('/api/articlelike/create/'+params.id, {
    method: 'POST',
    credentials: 'include',
  });
}

export async function removeCollection(values) {
  return request('/api/removeCollection', {
    params: values,
    method: 'POST',
  });
}

export async function createCollection(values) {
  return request('/api/removeCollection', {
    params: values,
    method: 'POST',
  });
}



export async function getRelation(values) {
  return request('/api/getRelation', {
  });
}

