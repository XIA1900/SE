import { Avatar } from 'antd';
import React from 'react';
import moment from 'moment';
import { history } from 'umi';
import styles from './index.less';

const ArticleListContent = ({data: { content, avatar, createdAt, name, group, groupID },}) => {

  const clickGroup = () => {
    history.push({
      pathname: '/group/content',
      search: groupID.toString(),
    });
  }

  const clickUser = async(values) => {
    history.push({
      pathname: '/user/account',
      search: values,
    });
  }

  return ( 
  <div className={styles.listContent}>
    <div className={styles.description}>{content}</div>
    <div className={styles.extra}>
      <img src={avatar} style={{ width: '25px', height: '25px', borderRadius: '25px' }} />
      <a onClick={(e) => clickUser(name, e)}> {name}</a> posted on
      <a onClick={clickGroup}> {group}</a>
      <em>{moment(createdAt).format('YYYY-MM-DD HH:mm')}</em>
    </div>
  </div>
)};

export default ArticleListContent;
