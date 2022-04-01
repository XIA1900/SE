describe('homepage', function(){
    this.beforeEach(()=>{
    cy.visit('http://localhost:8001/user/login')
    cy.get('.ant-pro-form-login-main').get('#username').type('kirby')
    cy.get('.ant-pro-form-login-main').get('#password').type('007')
    cy.get('button:contains("Login")').click()
    cy.url().should('include', '/homepage')
        })
    it('forwardGrouppage', function(){
    cy.get('a:contains("Balala")').click()
    cy.url().should('include', '/group/content?21')
    })
    it('forwardPostpage', function(){
        cy.get('.ant-space-item').get('.anticon anticon-edit').click()
        cy.url().should('include', '/form/basic-form')
        })
    it('forwardPersonalCenter', function(){
            cy.get('.ant-space-item').get('.ant-dropdown-trigger action__LP4_P account__6HXOq').click()
            cy.url().should('include', '/account/center?kirby')
        })
    it('forwardPersonalSettings', function(){
            cy.get('.ant-space-item').get('.ant-dropdown-trigger action__LP4_P account__6HXOq').click()
            cy.url().should('include', '/account/settings?kirby')
        }) 
        it('forwardCreatGroup', function(){
            cy.get('.ant-space-item').get('.ant-dropdown-trigger action__LP4_P account__6HXOq').click()
            cy.url().should('include', '/account/selectGroups/created?kirby')
        }) 
        it('Logout', function(){
            cy.get('.ant-space-item').get('.ant-dropdown-trigger action__LP4_P account__6HXOq').click()
            cy.url().should('include', '/user/login')
        })     
})
