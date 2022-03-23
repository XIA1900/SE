/*
A post have:
0. postid
1. title
2. content
3. owner(owner name, owner avatar)
4. last updated at
5. replies_count and replies (each reply has owner(name and avatar), likes, content, createdAt and replies)
6. likes_count and likes(users who like this post)
7. collections_count and collections(users who collect this post)
*/

/*
url: /group/post?postid
*/

import { PlusOutlined, TeamOutlined, CrownOutlined, CalendarOutlined } from '@ant-design/icons';
import { Avatar, Card, Col, Divider, Input, Row, Tag } from 'antd';
import React, { useState, useRef } from 'react';
import { GridContent } from '@ant-design/pro-layout';
import { Link, useRequest, history } from 'umi';
import Like from './components/like';
import Reply from './components/reply';
import Collection from './components/collection';
import styles from './Center.less';
import { getPost } from '@/services/getPost';

const postid = history.location.search.substring(1);
console.log(postid);

const operationTabList = [
  {
    key: 'reply',
    tab: <span> Replies </span>,
  },
  {
    key: 'like',
    tab: <span> Likes </span>,
  },
  {
    key: 'collection',
    tab: <span> Collections</span>
  },
];

const TagList = ({ tags }) => {
  const ref = useRef(null);
  const [newTags, setNewTags] = useState([]);
  const [inputVisible, setInputVisible] = useState(false);
  const [inputValue, setInputValue] = useState('');

  const showInput = () => {
    setInputVisible(true);

    if (ref.current) {
      // eslint-disable-next-line no-unused-expressions
      ref.current?.focus();
    }
  };

  const handleInputChange = (e) => {
    setInputValue(e.target.value);
  };

  const handleInputConfirm = () => {
    let tempsTags = [...newTags];

    if (inputValue && tempsTags.filter((tag) => tag.label === inputValue).length === 0) {
      tempsTags = [
        ...tempsTags,
        {
          key: `new-${tempsTags.length}`,
          label: inputValue,
        },
      ];
    }

    setNewTags(tempsTags);
    setInputVisible(false);
    setInputValue('');
  };

  return null;
};

const Center = () => {
  const [tabKey, setTabKey] = useState('reply');

  const { data: postContents, loading } = useRequest(() => {
    return getPost({
      postid: postid,
    });
  });

  const list = postContents?.list || [];
  console.log(list);

  const renderPostInfo = ({ avatar, title, content, owner, updatedAt }) => {
    return (
      <div className={styles.listContent}>
        <div className={styles.title}>{title}</div>
          <img
            alt=""
            src={avatar}
            style={{ width: '25px', height: '25px', borderRadius: '25px' }}
          />
          <a href=''> {owner}</a> updated at {updatedAt}
        <div className={styles.description}> {content} </div>
      </div>
    );
  };

  // 渲染tab切换

  const renderChildrenByTabKey = (tabValue) => {
    if (tabValue === 'reply') {
      return <Reply />;
    }

    if (tabValue === 'like') {
      return <Like />;
    }

    if (tabValue === 'collection') {
        return <Collection />;
    }

    return null;
  };

  return (
    <GridContent>
      <Row gutter={24}>
        <Col lg={17} md={24}>
          <Card
            bordered={false}
            style={{
              marginBottom: 0,
            }}
            loading={loading}
          >
            {!loading && list && (
              <div>
                {renderPostInfo(list)}
              </div>
            )}
          </Card>

          <Card
            className={styles.tabsCard}
            bordered={false}
            tabList={operationTabList}
            activeTabKey={tabKey}
            onTabChange={(_tabKey) => {
              setTabKey(_tabKey);
            }}
          >
            {renderChildrenByTabKey(tabKey)}
          </Card>
        </Col>
      </Row>
    </GridContent>
  );
};

export default Center;
