// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
contract Verification{

    enum DocStatus{
        accepted,
        rejected,
        pending
    }
    struct User{
        string  name;
        string  email;
    }

    struct Verifier{
        string  name;
        string  email;
        string  AadharNumber;
        string  institute;
        //Will need this in ./documents.sol to check if account verifying the documents
        //is actually a verifier. 
        bool    isVerifier;
    }
    struct Document{
        address     requester;
        address     verifiedBy; 
        string      name;
        string      description; 
        string      docAddressOnIPFS;
        uint        docIndex;
        DocStatus   status;     
    }


    mapping(address=>Verifier) public verifiers;
    mapping(address=>User) private users;
    mapping(string=>Document) private documentList;
    mapping(address=>bool) userList;
    mapping(address=>Document[]) userUploadedDocuments;

   
    Document[] documents;
        
    uint docIndexCounter=0;
    uint userDocIndexCounter=0;
    function registerAsUser(string calldata _name, string calldata _email) public{
        users[msg.sender]=User({
            name:       _name,
            email:      _email
        });
        userList[msg.sender]=true;
    }
    function registerAsVerifier(string memory _name, string memory _email, string memory _aadharNum, string memory _institute) public{
        verifiers[msg.sender]=Verifier({
            name:           _name,
            email:          _email,
            AadharNumber:   _aadharNum,
            institute:      _institute,
            isVerifier:     true
        });
    }



    function addDocument(string calldata _name, string calldata _description, string calldata _docAddressOnIPFS) public {
        bool isUser=userList[msg.sender];
        
        require(isUser==true,"register first to verify");

        Document memory docs=Document({
            requester:          msg.sender,
            verifiedBy:         address(0),
            name:               _name,
            description:        _description,
            docAddressOnIPFS:   _docAddressOnIPFS,
            docIndex:           docIndexCounter,
            status:             DocStatus.pending
        });


        documentList[_docAddressOnIPFS]=docs;
        documents.push(docs);
        docIndexCounter++;

        userUploadedDocuments[msg.sender].push(docs);
    }

    function getDocumentList() public view  returns (Document[] memory){
        if(verifiers[msg.sender].isVerifier){
            return (documents);
        }
        return (userUploadedDocuments[msg.sender]);
    }

    function verifyDocuments(string memory _docAddressOnIPFS, DocStatus _status) public payable {
        bool isVerifier=verifiers[msg.sender].isVerifier;
        require(isVerifier==true,"You're not verifier");
        documentList[_docAddressOnIPFS].status=_status;
        documentList[_docAddressOnIPFS].verifiedBy=msg.sender;
        uint index = documentList[_docAddressOnIPFS].docIndex;
        documents[index].status=_status;
        documents[index].verifiedBy=msg.sender;
    }
 
}
