import { PlusOutlined, HomeOutlined, ContactsOutlined, ClusterOutlined } from '@ant-design/icons';
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

  const { data, loading } = useRequest(() => {
    return getGroupBasic({
      groupName,
    });
  });

  const list = data?.list || [];
  console.log(list);

  const renderGroupInfo = ({ data }) => {
    return (
      <div className={styles.detail}>
        <h>{list.groupName}</h>
        <p>{list.groupDescription}</p>
        <p>
          <ContactsOutlined
            style={{
              marginRight: 8,
            }}
          />
          {data.list.groupOwner}
        </p>
        <p>
          <ClusterOutlined
            style={{
              marginRight: 8,
            }}
          />
          {data.list.groupMember}
        </p>
        <p>
          <HomeOutlined
            style={{
              marginRight: 8,
            }}
          />
          <p> Created at {data.list.createdAt} </p>
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
            {!loading && data && (
              <div>
                <div className={styles.avatarHolder}>
                  <img alt="" src={data.list.groupAvatar} />
                  {/* <div className={styles.name}>{currentUser.name}</div>
                  <div>{currentUser?.signature}</div> */}
                </div>
                {renderGroupInfo(data)}
                <Divider dashed />
                {/* <TagList tags={currentUser.tags || []} /> */}
                <Divider
                  style={{
                    marginTop: 16,
                  }}
                  dashed
                />
                {/* <div className={styles.team}>
                  <div className={styles.teamTitle}>团队</div>
                  <Row gutter={36}>
                    {currentUser.notice &&
                      currentUser.notice.map((item) => (
                        <Col key={item.id} lg={24} xl={12}>
                          <Link to={item.href}>
                            <Avatar size="small" src={item.logo} />
                            {item.member}
                          </Link>
                        </Col>
                      ))}
                  </Row>
                </div> */}
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
