import { PlusOutlined, TeamOutlined, CrownOutlined, CalendarOutlined, FormOutlined, FrownOutlined, SmileOutlined  } from '@ant-design/icons';
import { Button, Avatar, Card, Col, Divider, Input, Row, Tag } from 'antd';
import React, { useState, useRef } from 'react';
import { GridContent } from '@ant-design/pro-layout';
import { Link, useRequest, history, useModel } from 'umi';
import Hottest from './components/hottest';
import Latest from './components/latest';
import styles from './Center.less';
import { getGroupBasic } from '@/services/getGroupInfo';
import { checkMember, quitGroup, joinGroup } from '@/services/user';
import { countReset } from 'console';


const groupID = history.location.search.substring(1);
const pageNo = 1;
const pageSize = 10;

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
  const { initialState, setInitialState } = useModel('@@initialState');
  const { currentUser } = initialState;

  const { data, loading } = useRequest( async() => {
    const result = await getGroupBasic({
      groupID: groupID,
      username: currentUser.name,
      pageNO: pageNo,
      pageSize: pageSize,
    });
    return result;
    },
    {
      formatResult: result => result,
      loadMore: true,
    }
  );

  console.log(data);
  let list = [];
  if(typeof(data.community) != 'undefined') {
    const community = data.community;
    list = {
      id: community.ID,
      groupOwner: community.Creator,
      groupName: community.Name,
      groupDescription: community.Description, 
      createdAt: community.CreateDay, 
      groupMember: data.count,
      ifexit: data.ifexit,
    };
  }

  //console.log(list);
  
  //const list = groupBasics?.list || [];

  const isMember = async() => {
    return await checkMember({
      groupName: groupName,
      user: currentUser.name,
    });
  };

  const onJoin = async() => {
    const result = await joinGroup({
      groupName: groupName,
      user: currentUser.name
    });
    console.log(result);
    if(result.message === 'Ok') {
      location.reload();
    }
  };

  const onQuit = async() => {
    const result = await quitGroup({
      groupName: groupName,
      user: currentUser.name,
    });
    if(result.message === 'Ok') {
      location.reload();
    }
  };

  const onPost = async() => {
    history.push({
      pathname: '/form/createPost',
      search: groupName,
    })
  };

  const renderGroupInfo = ({ groupOwner, groupName, groupDescription, createdAt, groupMember, ifexit }) => {
    if(ifexit === true) {
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
          <Button onClick={onQuit} style={{display: 'inline-block'}}>
            <FrownOutlined/>
              Quit
          </Button> 
          &nbsp;&nbsp;&nbsp;&nbsp;
          <Button onClick={onPost} style={{display: 'inline-block'}}>
            <FormOutlined />
            Post
          </Button>
        </div>
      );
    }
    else {
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
          <Button onClick={onJoin}>
            <SmileOutlined/>
              Join
          </Button>
        </div>
      );
    }
    
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
