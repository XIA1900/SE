import { Avatar } from 'antd';
import React from 'react';
import moment from 'moment';
import styles from './index.less';
import { group } from 'console';

const ArticleListContent = ({
  data: { Content, logo, CreateDay, Username, CommunityID },
  
}) => {
  const group_href = '/group/content?' + CommunityID;
  return ( 
  <div className={styles.listContent}>
    <div className={styles.description}>{Content}</div>
    <div className={styles.extra}>
      <img src={logo} style={{ width: '25px', height: '25px', borderRadius: '25px' }} />
      <a href=''> {Username}</a> posted on
      <a href={group_href}> {CommunityID}</a>
      <em>{moment(CreateDay).format('YYYY-MM-DD HH:mm')}</em>
    </div>
  </div>
)};

export default ArticleListContent;
