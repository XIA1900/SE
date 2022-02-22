import { request } from 'umi';


/*
search something: 

input: values is a string

return: most related 10 articles, same properties as /api/queryList
*/

export async function search(values) {
  return request('/api/search', {
    values,
  });
}
