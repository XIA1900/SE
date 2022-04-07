import { request } from 'umi';

export async function createGroup(params) {
  return request('/api/createGroup', {
    data: params,
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
  });
}

export async function createPost(body) {
  console.log(body);
  return request('/api/article/create', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    credentials: 'include',
  });
}
