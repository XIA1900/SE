import { request } from 'umi';

export async function getGroupInfo(userName) {   //only for created groups, return entire information
    return request('/api/getGroupInfo', {
      userName,
    });
  }

export async function getCreatedGroup(userName) {  //only return group basic information, number of member,number of lists
  return request('/api/getCreatedGroup', {
    userName,
  });
}

export async function getJoinedGroup(userName) {  //only return group basic information, group link
  return request('api/getJoinedGroup', {
    userName,
  });
}