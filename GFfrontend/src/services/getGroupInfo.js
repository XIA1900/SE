import { request } from 'umi';

export async function getGroupPosts(params) {
  //only for created groups, return entire information
  return request('/api/article/getarticlelistbycommunityid?CommunityID='+params.id+'&PageNO='+params.pageNO+'&PageSize='+params.pageSize, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
    credentials: 'include',
  });
}

export async function getCreatedGroup(params) {
  //only return group basic information, number of member,number of lists
  return request('/api/getCreatedGroup', {
    params,
  });
}

export async function getJoinedGroup(params) {
  //only return group basic information, group link
  return request('/api/getJoinedGroup', {
    params,
  });
}

export async function getGroupBasic(params) {
  //only for created groups, return entire information
  return request('/api/community/getone?id='+params.groupID+'&username='+params.username+'&pageNO='+params.pageNO+'&pageSize='+params.pageSize, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
    credentials: 'include',
  });
}