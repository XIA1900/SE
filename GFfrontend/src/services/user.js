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
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
    credentials: 'include',
  });
}

export async function getPersonalCollection(values) {
  return request('/api/getPersonalCollection');
}

export async function removeCollection(values) {
  return request('/api/removeCollection', {
    params: values,
    method: 'POST',
  });
}


export async function getPersnalFollower(values) {
  return request('/api/getPersnalFollower');
}

export async function removeFollower(values) {
  return request('/api/removeFollower', {
    params: values,
    method: 'POST',
  });
}

export async function getPersonalFollowing(values) {
  return request('/api/getPersonalFollowing');
}
  
export async function removeFollowing(values) {
  return request('/api/removeFollowing', {
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

export async function getRelation(values) {
  return request('/api/getRelation', {
  });
}