import { request } from 'umi';

export async function getGroupInfo(params) {   //only for created groups, return entire information
    return request('/api/getGroupInfo', {
      params,
    });
  }

export async function getCreatedGroup(params) {  //only return group basic information, number of member,number of lists
  return request('/api/getCreatedGroup', {
    params,
  });
}

export async function getJoinedGroup(params) {  //only return group basic information, group link
  return request('/api/getJoinedGroup', {
    params,
  });
}