import { CalendarOutlined, PlusOutlined, HomeOutlined, ContactsOutlined, ClusterOutlined, PhoneOutlined, MailOutlined, WomanOutlined, ManOutlined } from '@ant-design/icons';
import { Avatar, Card, Col, Divider, Input, Row, Tag } from 'antd';
import React, { useState, useRef } from 'react';
import { GridContent } from '@ant-design/pro-layout';
import { Link, useRequest, history } from 'umi';
import Follower from './components/Follower';
import Following from './components/Following';
import { queryCurrent } from '@/services/user';
import styles from './Center.less';
import { domainToASCII } from 'url';

/*other user's view of a user*/

const username = history.location.search.substring(1);

const operationTabList = [
  {
    key: 'follower',
    tab: (
      <span>
        Follower{' '}
        <span
          style={{
            fontSize: 14,
          }}
        >
        </span>
      </span>
    ),
  },
  {
    key: 'following',
    tab: (
      <span>
        Following{' '}
        <span
          style={{
            fontSize: 14,
          }}
        >
        </span>
      </span>
    ),
  },
];

const CourseList = ({ tags }) => {
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

  return (
    <div className={styles.tags}>
      <div className={styles.tagsTitle}> Courses </div>
      {(tags || []).concat(newTags).map((item) => (
        <Tag key={item.key}>{item.label}</Tag>
      ))}
      {inputVisible && (
        <Input
          ref={ref}
          type="text"
          size="small"
          style={{
            width: 78,
          }}
          value={inputValue}
          onChange={handleInputChange}
          onBlur={handleInputConfirm}
          onPressEnter={handleInputConfirm}
        />
      )}
      {!inputVisible && (
        <Tag
          onClick={showInput}
          style={{
            borderStyle: 'dashed',
          }}
        >
          <PlusOutlined />
        </Tag>
      )}
    </div>
  );
};

const InterestList = ({ tags }) => {
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

  return (
    <div className={styles.tags}>
      <div className={styles.tagsTitle}> Interests </div>
      {(tags || []).concat(newTags).map((item) => (
        <Tag key={item.key}>{item.label}</Tag>
      ))}
      {inputVisible && (
        <Input
          ref={ref}
          type="text"
          size="small"
          style={{
            width: 78,
          }}
          value={inputValue}
          onChange={handleInputChange}
          onBlur={handleInputConfirm}
          onPressEnter={handleInputConfirm}
        />
      )}
      {!inputVisible && (
        <Tag
          onClick={showInput}
          style={{
            borderStyle: 'dashed',
          }}
        >
          <PlusOutlined />
        </Tag>
      )}
    </div>
  );
};

const Center = () => {
  const [tabKey, setTabKey] = useState('follower'); //  获取用户信息
  const { initialState, setInitialState } = useModel('@@initialState');
  const { currentUser } = initialState;

  const { data, loading } = useRequest(
    async() => {
      const result = await queryCurrent({
        username: username,
      });
      return result;
    },
    {
      formatResult: result => result,
    }
  ); //  渲染用户信息
  
  console.log(data);
  let visitedUser = [];
  if(typeof(data) != 'undefined') {
    visitedUser = {
      name: data.Username,
      birthday: data.Birthday,
      gender: data.Gender,
      major: data.Department,
      avatar: 'http://192.168.3.132:10010/resources/userfiles/'+ data.Username+'/avatar.png',
      follower: true,  //current user is a follower of visited user
      mutual: true,
      blacklist: false, //visited user is in blacklist of current user
    };
  }

  const unfollow = () => {

  }

  const follow = () => {

  }

  const renderButton = ({follower, mutual}) => {
    if(mutual === true) {
      return (
        <div>
          <Button onClick={unfollow}>
          Mutual
          </Button>
          <Button onClick={onBlock}>
            Block
          </Button> 
        </div>
      );
    }
    if(follower === true) {
      return (
        <div>
          <Button onClick={unfollow}>
            Following
          </Button>
          <Button onClick={onBlock}>
            Block
          </Button> 
        </div>
        
      )
    }
    else {
      return (
        <div>
          <Button onClick={follow}>
            Follow
          </Button>
          <Button onClick={onBlock}>
            Block
          </Button> 
        </div>

      )
    }
  }

  const renderGender = ({gender}) => {
    if(gender === 'Female') {
      return (
        <WomanOutlined/>
      );
    }
    else if(gender === 'Male') {
      return (
        <ManOutlined/>
      );
    }
    else {
      return;
    }
  }

  const renderUserInfo = ({ birthday, gender, email, major, grade, country, province, city, phone }) => {
    
      return (
        <div className={styles.detail}>
          <p>
            <CalendarOutlined
              style={{
                marginRight: 8,
              }}
            />
            {birthday+'    '}
            {renderGender(gender)}
          </p>
          
          <p>
            <MailOutlined
              style={{
                marginRight: 8,
              }}
            />
            {email}
          </p>
          <p>
            <PhoneOutlined
              style={{
                marginRight: 8,
              }}
            />
            {phone} 
          </p>
          <p>
            <ClusterOutlined
              style={{
                marginRight: 8,
              }}
            />
            {major+' '}{grade} 
          </p>
          <p>
            <HomeOutlined
              style={{
                marginRight: 8,
              }}
            />
            {country+' '}{province+' '}{city}
          </p>
        </div>
      );
  }; // 渲染tab切换

  // const renderUserInfo = ({ birthday, gender, email, major, grade, country, province, city, phone }) => {
  //   if(gender === 'Female') {
  //     return (
  //       <div className={styles.detail}>
  //         <p>
  //           <CalendarOutlined
  //             style={{
  //               marginRight: 8,
  //             }}
  //           />
  //           {birthday+'    '}
  //           <WomanOutlined/>
  //         </p>
          
  //         <p>
  //           <MailOutlined
  //             style={{
  //               marginRight: 8,
  //             }}
  //           />
  //           {email}
  //         </p>
  //         <p>
  //           <PhoneOutlined
  //             style={{
  //               marginRight: 8,
  //             }}
  //           />
  //           {phone} 
  //         </p>
  //         <p>
  //           <ClusterOutlined
  //             style={{
  //               marginRight: 8,
  //             }}
  //           />
  //           {major+' '}{grade} 
  //         </p>
  //         <p>
  //           <HomeOutlined
  //             style={{
  //               marginRight: 8,
  //             }}
  //           />
  //           {country+' '}{province+' '}{city}
  //         </p>
  //       </div>
  //     );
  //   }
  //   else if (gender === 'Male') {
  //     return (
  //       <div className={styles.detail}>
  //         <p>
  //           <CalendarOutlined
  //             style={{
  //               marginRight: 8,
  //             }}
  //           />
  //           {birthday+'    '} 
  //           <ManOutlined/>
  //         </p>
          
  //         <p>
  //           <MailOutlined
  //             style={{
  //               marginRight: 8,
  //             }}
  //           />
  //           {email}
  //         </p>
  //         <p>
  //           <PhoneOutlined
  //             style={{
  //               marginRight: 8,
  //             }}
  //           />
  //           {phone} 
  //         </p>
  //         <p>
  //           <ClusterOutlined
  //             style={{
  //               marginRight: 8,
  //             }}
  //           />
  //           {major+' '}{grade} 
  //         </p>
  //         <p>
  //           <HomeOutlined
  //             style={{
  //               marginRight: 8,
  //             }}
  //           />
  //           {country+' '}{province+' '}{city}
  //         </p>
  //       </div>
  //     );
  //   }
  //   else {
  //     return (
  //       <div className={styles.detail}>
  //         <p>
  //           <CalendarOutlined
  //             style={{
  //               marginRight: 8,
  //             }}
  //           />
  //           {birthday+'    '} 
  //         </p>
          
  //         <p>
  //           <MailOutlined
  //             style={{
  //               marginRight: 8,
  //             }}
  //           />
  //           {email}
  //         </p>
  //         <p>
  //           <PhoneOutlined
  //             style={{
  //               marginRight: 8,
  //             }}
  //           />
  //           {phone} 
  //         </p>
  //         <p>
  //           <ClusterOutlined
  //             style={{
  //               marginRight: 8,
  //             }}
  //           />
  //           {major+' '}{grade} 
  //         </p>
  //         <p>
  //           <HomeOutlined
  //             style={{
  //               marginRight: 8,
  //             }}
  //           />
  //           {country+' '}{province+' '}{city}
  //         </p>
  //       </div>
  //     );
  //   }
  // }; // 渲染tab切换

  const renderChildrenByTabKey = (tabValue) => {
    if (tabValue === 'follower') {
      return <Follower />;
    }

    if (tabValue === 'following') {
      return <Following />;
    }

    return null;
  };

  return (
    <GridContent>
      <Row gutter={24}>
        <Col lg={7} md={24}>
          <Card
            bordered={false}
            style={{
              marginBottom: 24,
            }}
            loading={loading}
          >
            {!loading && visitedUser && (
              <div>
                <div className={styles.avatarHolder}>
                  <img alt="" src={visitedUser.avatar} />
                  <div className={styles.name}>{visitedUser.name}</div>
                  <div>{visitedUser?.signature}</div>
                </div>
                {renderButton(visitedUser)}
                {renderUserInfo(visitedUser)}
                <Divider dashed />
                <CourseList tags={visitedUser.courses || []} />
                <Divider
                  style={{
                    marginTop: 16,
                  }}
                  dashed
                />
                <InterestList tags={visitedUser.interests || []} />
              </div>
            )}
          </Card>
        </Col>
        <Col lg={17} md={24}>
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
