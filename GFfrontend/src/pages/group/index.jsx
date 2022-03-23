import { PlusOutlined, TeamOutlined, CrownOutlined, CalendarOutlined } from '@ant-design/icons';
import { Avatar, Card, Col, Divider, Input, Row, Tag } from 'antd';
import React, { useState, useRef } from 'react';
import { GridContent } from '@ant-design/pro-layout';
import { Link, useRequest, history } from 'umi';
import Hottest from './components/hottest';
import Latest from './components/latest';
import styles from './Center.less';
import { getGroupBasic } from '@/services/getGroupInfo';
const groupName = history.location.search.substring(1);
const operationTabList = [
  {
    key: 'hottest',
    tab: <span>Hottest </span>,
  },
  {
    key: 'latest',
    tab: <span>Latest </span>,
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
  const [tabKey, setTabKey] = useState('hottest');

  const { data: groupBasics, loading } = useRequest(() => {
    return getGroupBasic({
      groupName,
    });
  });

  const list = groupBasics?.list || [];
  console.log('groups');
  console.log(list);

  const renderGroupInfo = ({ groupOwner, groupName, groupDescription, createdAt, groupMember }) => {
    return (
      <div className={styles.detail}>
        <h1>{groupName}</h1>
        <p>{groupDescription}</p>
        <p>
          <CrownOutlined
            style={{
              marginRight: 8,
            }}
          />
          {groupOwner}
          <TeamOutlined
            style={{
              marginRight: 8,
              marginLeft: 20,
            }}
          />
          {groupMember}
          <CalendarOutlined
            style={{
              marginRight: 8,
              marginLeft: 20,
            }}
          />
          Created at {createdAt}
        </p>
      </div>
    );
  };

  // 渲染tab切换

  const renderChildrenByTabKey = (tabValue) => {
    if (tabValue === 'hottest') {
      return <Hottest />;
    }

    if (tabValue === 'latest') {
      return <Latest />;
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
              marginBottom: 24,
            }}
            loading={loading}
          >
            {!loading && list && (
              // <div>
              <div className={styles.avatarHolder}>
                <img
                  alt=""
                  src={list.groupAvatar}
                  style={{ width: '100px', height: '100px', borderRadius: '100px' }}
                />
                {renderGroupInfo(list)}
              </div>
              // {/* </div> */}
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
