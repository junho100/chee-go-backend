package service

import (
	"chee-go-backend/internal/domain/entity"
	"chee-go-backend/internal/domain/repository"
	"chee-go-backend/internal/domain/service"
	"chee-go-backend/internal/http/dto"
	"time"
)

type resumeService struct {
	resumeRepository repository.ResumeRepository
}

func NewResumeService(resumeRepository repository.ResumeRepository) service.ResumeService {
	return &resumeService{
		resumeRepository: resumeRepository,
	}
}

func (r *resumeService) ConvertResumeToLinkedin(resume entity.Resume, keywords []string) dto.LinkedinResume {
	linkedinResume := &dto.LinkedinResume{
		Introduction:    resume.Introduction,
		WorkExperiences: make([]dto.LinkedinResumeWorkExperience, len(resume.WorkExperiences)),
		Educations:      make([]dto.LinkedinResumeEducation, len(resume.Educations)),
		Certificates:    make([]dto.LinkedinResumeCertificate, len(resume.Certificates)),
		Projects:        make([]dto.LinkedinResumeProject, len(resume.Projects)),
		Skills:          keywords,
	}

	for i, workExperience := range resume.WorkExperiences {
		linkedinResume.WorkExperiences[i].Position = workExperience.Position
		linkedinResume.WorkExperiences[i].CompanyName = workExperience.CompanyName
		linkedinResume.WorkExperiences[i].StartDate = workExperience.StartDate
		linkedinResume.WorkExperiences[i].EndDate = workExperience.EndDate

		var content string
		for _, detail := range workExperience.WorkExperienceDetails {
			content += (detail.Name + "\n")
			content += (detail.Content + "\n\n")
		}
		linkedinResume.WorkExperiences[i].Content = content
	}

	for i, education := range resume.Educations {
		linkedinResume.Educations[i].SchoolName = education.SchoolName
		linkedinResume.Educations[i].MajorName = education.MajorName
		linkedinResume.Educations[i].StartDate = education.StartDate
		linkedinResume.Educations[i].EndDate = education.EndDate
	}

	for i, certificate := range resume.Certificates {
		linkedinResume.Certificates[i].Name = certificate.Name
		linkedinResume.Certificates[i].IssuedBy = certificate.IssuedBy
		linkedinResume.Certificates[i].IssuedDate = certificate.IssuedDate
	}

	for i, project := range resume.Projects {
		linkedinResume.Projects[i].Name = project.Name
		linkedinResume.Projects[i].StartDate = project.StartDate
		linkedinResume.Projects[i].EndDate = project.EndDate

		content := (project.Summary + "\n\n" + project.Content)
		content += ("\n\n" + "Github: " + project.GithubURL)
		linkedinResume.Projects[i].Content = content
	}

	return *linkedinResume
}

func (r *resumeService) ConvertResumeToProgrammers(resume entity.Resume) dto.ProgrammersResume {
	programmersResume := &dto.ProgrammersResume{
		WorkExperiences: make([]dto.ProgrammersResumeWorkExperience, len(resume.WorkExperiences)),
		Educations:      make([]dto.ProgrammersResumeEducation, len(resume.Educations)),
		Projects:        make([]dto.ProgrammersResumeProject, len(resume.Projects)),
		Certificates:    make([]dto.ProgrammersResumeCertificates, len(resume.Certificates)),
		Activities:      make([]dto.ProgrammersResumeActivity, len(resume.Activities)),
	}

	for i, workExperience := range resume.WorkExperiences {
		programmersResume.WorkExperiences[i].CompanyName = workExperience.CompanyName
		programmersResume.WorkExperiences[i].Position = workExperience.Position
		programmersResume.WorkExperiences[i].StartDate = workExperience.StartDate

		programmersResume.WorkExperiences[i].Details = make([]dto.ProgrammersResumeWorkExperienceDetail, len(workExperience.WorkExperienceDetails))
		for j, detail := range resume.WorkExperiences[i].WorkExperienceDetails {
			programmersResume.WorkExperiences[i].Details[j].Name = detail.Name
			programmersResume.WorkExperiences[i].Details[j].StartDate = detail.StartDate
			programmersResume.WorkExperiences[i].Details[j].EndDate = detail.EndDate
			programmersResume.WorkExperiences[i].Details[j].Content = detail.Content
		}
	}

	for i, education := range resume.Educations {
		programmersResume.Educations[i].SchoolName = education.SchoolName
		programmersResume.Educations[i].MajorName = education.MajorName
		programmersResume.Educations[i].StartDate = education.StartDate
		programmersResume.Educations[i].EndDate = education.EndDate
	}

	for i, project := range resume.Projects {
		programmersResume.Projects[i].Name = project.Name
		programmersResume.Projects[i].StartDate = project.StartDate
		programmersResume.Projects[i].Summary = project.Summary
		programmersResume.Projects[i].Content = project.Content
		programmersResume.Projects[i].GithubURL = project.GithubURL
	}

	for i, certificate := range resume.Certificates {
		programmersResume.Certificates[i].Name = certificate.Name
		programmersResume.Certificates[i].IssuedBy = certificate.IssuedBy
		programmersResume.Certificates[i].IssuedDate = certificate.IssuedDate
	}

	for i, activity := range resume.Activities {
		programmersResume.Activities[i].Name = activity.Name
		programmersResume.Activities[i].StartDate = activity.StartDate
		programmersResume.Activities[i].EndDate = activity.EndDate
		programmersResume.Activities[i].Content = activity.Content
	}

	return *programmersResume
}

func (r *resumeService) ConvertResumeToWanted(resume entity.Resume, keywords []string) dto.WantedResume {
	wantedResume := &dto.WantedResume{
		Introduction:    resume.Introduction,
		WorkExperiences: make([]dto.WantedResumeWorkExperience, len(resume.WorkExperiences)),
		Educations:      make([]dto.WantedResumeEducation, len(resume.Educations)),
		Certificates:    make([]dto.WantedResumeCertificate, len(resume.Certificates)+len(resume.Projects)+len(resume.Activities)),
		Skills:          keywords,
	}

	for i, workExperience := range resume.WorkExperiences {
		wantedResume.WorkExperiences[i] = dto.WantedResumeWorkExperience{
			CompanyName: workExperience.CompanyName,
			Position:    workExperience.Position,
			StartDate:   workExperience.StartDate,
			EndDate:     workExperience.EndDate,
		}

		wantedResume.WorkExperiences[i].Details = make([]dto.WantedResumeWorkExperienceDetail, len(resume.WorkExperiences[i].WorkExperienceDetails))
		for j, detail := range resume.WorkExperiences[i].WorkExperienceDetails {
			wantedResume.WorkExperiences[i].Details[j] = dto.WantedResumeWorkExperienceDetail{
				Name:      detail.Name,
				StartDate: detail.StartDate,
				EndDate:   detail.EndDate,
				Content:   detail.Content,
			}
		}
	}

	for i, education := range resume.Educations {
		wantedResume.Educations[i] = dto.WantedResumeEducation{
			SchoolName: education.SchoolName,
			MajorName:  education.MajorName,
			StartDate:  education.StartDate,
			EndDate:    education.EndDate,
		}
	}

	wantedResume.Certificates = make([]dto.WantedResumeCertificate, len(resume.Certificates)+len(resume.Projects)+len(resume.Activities))
	idx := 0
	for _, certificate := range resume.Certificates {
		wantedResume.Certificates[idx].Name = certificate.Name
		wantedResume.Certificates[idx].StartDate = certificate.IssuedDate
		wantedResume.Certificates[idx].Content = certificate.IssuedBy
		idx++
	}

	for _, project := range resume.Projects {
		wantedResume.Certificates[idx].Name = project.Name
		wantedResume.Certificates[idx].StartDate = project.StartDate
		wantedResume.Certificates[idx].Content = project.Content
		idx++
	}

	for _, activity := range resume.Activities {
		wantedResume.Certificates[idx].Name = activity.Name
		wantedResume.Certificates[idx].StartDate = activity.StartDate
		wantedResume.Certificates[idx].Content = activity.Content
		idx++
	}

	return *wantedResume
}

func (s *resumeService) CreateResume(createResumeDto *dto.CreateResumeDTO) (uint, error) {
	resume := &entity.Resume{}
	var err error

	tx, err := s.resumeRepository.StartTransaction()
	if err != nil {
		return 0, nil
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	resume, err = s.resumeRepository.FindResumeByUserId(createResumeDto.UserID)

	if err != nil {
		resume = &entity.Resume{
			Introduction: createResumeDto.Introduction,
			GithubURL:    createResumeDto.GithubURL,
			BlogURL:      createResumeDto.BlogURL,
			UserID:       createResumeDto.UserID,
		}
	} else {
		resume.Introduction = createResumeDto.Introduction
		resume.GithubURL = createResumeDto.GithubURL
		resume.BlogURL = createResumeDto.BlogURL
		resume.UserID = createResumeDto.UserID
	}

	if err := s.resumeRepository.CreateResume(tx, resume); err != nil {
		tx.Rollback()
		return 0, err
	}

	if err := s.resumeRepository.DeleteEducationsInResume(tx, resume); err != nil {
		tx.Rollback()
		return 0, err
	}
	for _, education := range createResumeDto.Educations {
		savedEducation := &entity.Education{
			SchoolName: education.SchoolName,
			MajorName:  education.MajorName,
			ResumeID:   resume.ID,
		}

		if education.StartDate.IsZero() {
			savedEducation.StartDate = time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
		} else {
			savedEducation.StartDate = education.StartDate
		}

		if education.EndDate.IsZero() {
			savedEducation.EndDate = time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
		} else {
			savedEducation.EndDate = education.EndDate
		}

		if err := s.resumeRepository.CreateEducation(tx, savedEducation); err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	if err := s.resumeRepository.DeleteProjectsInResume(tx, resume); err != nil {
		tx.Rollback()
		return 0, err
	}
	for _, project := range createResumeDto.Projects {
		savedProject := &entity.Project{
			Content:   project.Content,
			GithubURL: project.GithubURL,
			ResumeID:  resume.ID,
			Name:      project.Name,
			Summary:   project.Summary,
		}

		if project.StartDate.IsZero() {
			savedProject.StartDate = time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
		} else {
			savedProject.StartDate = project.StartDate
		}

		if project.EndDate.IsZero() {
			savedProject.EndDate = time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
		} else {
			savedProject.EndDate = project.EndDate
		}

		if err := s.resumeRepository.CreateProject(tx, savedProject); err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	if err := s.resumeRepository.DeleteActivitiesInResume(tx, resume); err != nil {
		tx.Rollback()
		return 0, err
	}
	for _, activity := range createResumeDto.Activities {
		savedActivity := &entity.Activity{
			Name:     activity.Name,
			Content:  activity.Content,
			ResumeID: resume.ID,
		}

		if activity.StartDate.IsZero() {
			savedActivity.StartDate = time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
		} else {
			savedActivity.StartDate = activity.StartDate
		}

		if activity.EndDate.IsZero() {
			savedActivity.EndDate = time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
		} else {
			savedActivity.EndDate = activity.EndDate
		}

		if err := s.resumeRepository.CreateActivity(tx, savedActivity); err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	if err := s.resumeRepository.DeleteWorkExperiencesInResume(tx, resume); err != nil {
		tx.Rollback()
		return 0, err
	}
	for _, workExperience := range createResumeDto.WorkExperiences {
		savedWorkExperience := &entity.WorkExperience{
			CompanyName: workExperience.CompanyName,
			Department:  workExperience.Department,
			Position:    workExperience.Position,
			Job:         workExperience.Job,
			ResumeID:    resume.ID,
		}

		if workExperience.StartDate.IsZero() {
			savedWorkExperience.StartDate = time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
		} else {
			savedWorkExperience.StartDate = workExperience.StartDate
		}

		if workExperience.EndDate.IsZero() {
			savedWorkExperience.EndDate = time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
		} else {
			savedWorkExperience.EndDate = workExperience.EndDate
		}

		if err := s.resumeRepository.CreateWorkExperience(tx, savedWorkExperience); err != nil {
			tx.Rollback()
			return 0, err
		}

		for _, detail := range workExperience.Details {
			savedDetail := &entity.WorkExperienceDetail{
				Name:             detail.Name,
				Content:          detail.Content,
				WorkExperienceID: savedWorkExperience.ID,
			}

			if detail.StartDate.IsZero() {
				savedDetail.StartDate = time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
			} else {
				savedDetail.StartDate = detail.StartDate
			}

			if detail.EndDate.IsZero() {
				savedDetail.EndDate = time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
			} else {
				savedDetail.EndDate = detail.EndDate
			}

			if err := s.resumeRepository.CreateWorkExperienceDetail(tx, savedDetail); err != nil {
				tx.Rollback()
				return 0, err
			}
		}
	}

	if err := s.resumeRepository.DeleteKeywordsInResume(tx, resume); err != nil {
		tx.Rollback()
		return 0, err
	}
	for _, keywordName := range createResumeDto.Keywords {
		savedKeyword := &entity.Keyword{}
		if err := s.resumeRepository.FindKeywordByName(keywordName, savedKeyword); err != nil {
			if err := s.resumeRepository.CreateKeyword(tx, savedKeyword); err != nil {
				tx.Rollback()
				return 0, err
			}
		}

		savedKeywordResume := &entity.KeywordResume{
			ResumeID:  resume.ID,
			KeywordID: savedKeyword.ID,
		}
		if err := s.resumeRepository.CreateKeywordResume(tx, savedKeywordResume); err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	return resume.ID, tx.Commit().Error
}

func (r *resumeService) GetKeywordsByResumeID(id uint) []string {
	var keywords []entity.Keyword
	var err error
	keywords, err = r.resumeRepository.FindKeywordsByResumeId(id)
	if err != nil {
		return make([]string, 0)
	}

	keywordStrings := make([]string, len(keywords))
	for idx, keyword := range keywords {
		keywordStrings[idx] = keyword.Name
	}

	return keywordStrings
}

func (r *resumeService) GetResumeByUserID(userID string) (*entity.Resume, error) {
	resume, err := r.resumeRepository.FindResumeByUserId(userID)

	if err != nil {
		return nil, err
	}

	return resume, nil
}
