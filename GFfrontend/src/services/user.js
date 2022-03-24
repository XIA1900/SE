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