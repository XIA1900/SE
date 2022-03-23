import { Avatar } from 'antd';
import React from 'react';
import moment from 'moment';
import styles from './index.less';
import { group } from 'console';

const ArticleListContent = ({
  data: { content, logo, updatedAt, owner, group },
  
}) => {
  const group_href = '/group/content?' + group;
  return ( 
  <div className={styles.listContent}>
    <div className={styles.description}>{content}</div>
    <div className={styles.extra}>
      <img src={logo} style={{ width: '25px', height: '25px', borderRadius: '25px' }} />
      <a href=''> {owner}</a> posted on
      <a href={group_href}> {group}</a>
      <em>{moment(updatedAt).format('YYYY-MM-DD HH:mm')}</em>
    </div>
  </div>
)};

export default ArticleListContent;
